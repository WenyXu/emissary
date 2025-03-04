EMISSARY_CHART = $(OSS_HOME)/charts/emissary-ingress
YQ := $(OSS_HOME)/.circleci/yq
thisdir := $(patsubst %/,%,$(dir $(lastword $(MAKEFILE_LIST))))


define _push_chart
	CHART_NAME=$(1) $(OSS_HOME)/charts/scripts/push_chart.sh
endef

define _set_tag_and_repo
	$(OSS_HOME)/venv/bin/python $(OSS_HOME)/charts/scripts/update_chart_image_values.py \
		--values-file $(1) --tag $(2) --repo $(3)
endef

define _set_tag
	$(OSS_HOME)/venv/bin/python $(OSS_HOME)/charts/scripts/update_chart_image_values.py \
		--values-file $(1) --tag $(2)
endef

define _docgen
	if [[ -f $(1)/doc.yaml ]] ; then \
		$(thisdir)/docgen -d $(1)/doc.yaml -t $(1)/readme.tpl -v $(1)/values.yaml > $(1)/README.md ; \
	fi
endef

$(thisdir)/docgen: $(thisdir)/chart-doc-gen.d/go.mod
	cd $(<D) && go build -o $(abspath $@) kubepack.dev/chart-doc-gen

push-preflight: create-venv $(YQ)
	@$(OSS_HOME)/venv/bin/python -m pip install ruamel.yaml
.PHONY: push-preflight

release/ga/chart-push:
	for chart in $(EMISSARY_CHART) ; do \
		$(call _push_chart,`basename $$chart`) ; \
	done ;
.PHONY: release/ga/chart-push

chart-push-ci: push-preflight
	@([ $(IS_PRIVATE) ] && (echo "Private repo, not pushing chart" && exit 1)) || true
	@echo ">>> This will dirty your local tree and should only be run in CI"
	@echo ">>> If running locally, you'll probably want to run make chart-clean after running this"
	@[ -n "${CHART_VERSION_SUFFIX}" ] || (echo "CHART_VERSION_SUFFIX must be set for non-GA pushes" && exit 1)
	@[ -n "${IMAGE_TAG}" ] || (echo "IMAGE_TAG must be set" && exit 1)
	@[ -n "${IMAGE_REPO}" ] || (echo "IMAGE_REPO must be set" && exit 1)
	set -e; { \
		for chart in $(EMISSARY_CHART) ; do \
			sed -i.bak -E "s/version: ([0-9]+\.[0-9]+\.[0-9]+).*/version: \1${CHART_VERSION_SUFFIX}/g" $$chart/Chart.yaml && rm $$chart/Chart.yaml.bak ; \
			$(call _set_tag_and_repo,$$chart/values.yaml,${IMAGE_TAG},${IMAGE_REPO}) ; \
			$(YQ) w -i $$chart/Chart.yaml 'appVersion' ${IMAGE_TAG} ; \
			$(call _push_chart,`basename $$chart`) ; \
		done ; \
	}
.PHONY: chart-push-ci

release/chart/tag:
	@set -e; { \
		if [ -n "$(IS_DIRTY)" ]; then \
			echo "release/chart/tag: tree must be clean" >&2 ;\
			exit 1 ;\
		fi; \
		chart_ver=`grep 'version:' $(EMISSARY_CHART)/Chart.yaml | awk ' { print $$2 }'` ; \
		chart_ver=chart-v$${chart_ver} ; \
		git tag -m "Tagging $${chart_ver}" -a $${chart_ver} ; \
		git push origin $${chart_ver} ; \
	}

release/changelog:
	@for chart in $(EMISSARY_CHART) ; do \
		CHART_NAME=`basename $$chart` $(OSS_HOME)/charts/scripts/update_chart_changelog.sh ; \
	done ;
.PHONY: release/changelog

release/chart/update-images: $(YQ) $(thisdir)/docgen
	@[ -n "${IMAGE_TAG}" ] || (echo "IMAGE_TAG must be set" && exit 1)
	([[ "${IMAGE_TAG}" =~ .*\.0$$ ]] && $(MAKE) release/chart-bump/minor) || $(MAKE) release/chart-bump/revision
	for chart in $(EMISSARY_CHART) ; do \
		[[ "${IMAGE_TAG}" =~ .*\-ea$$ ]] && sed -i.bak -E "s/version: ([0-9]+\.[0-9]+\.[0-9]+).*/version: \1-ea/g" $$chart/Chart.yaml && rm $$chart/Chart.yaml.bak ; \
		$(call _set_tag,$$chart/values.yaml,${IMAGE_TAG}) ; \
		$(YQ) w -i $$chart/Chart.yaml 'appVersion' ${IMAGE_TAG} ; \
		IMAGE_TAG="${IMAGE_TAG}" CHART_NAME=`basename $$chart` $(OSS_HOME)/charts/scripts/image_tag_changelog_update.sh ; \
		CHART_NAME=`basename $$chart` $(OSS_HOME)/charts/scripts/update_chart_changelog.sh ; \
		$(call _docgen,$$chart) ; \
	done ;

chart/docgen: $(thisdir)/docgen
	for chart in $(EMISSARY_CHART) ; do \
		$(call _docgen,$$chart) ; \
	done ;
.PHONY: chart/docgen

# Both charts should have same versions for now. Just makes things a bit easier if we publish them together for now
release/chart-bump/revision:
	@for chart in $(EMISSARY_CHART) ; do \
		$(OSS_HOME)/charts/scripts/bump_chart_version.sh patch $$chart/Chart.yaml ; \
	done ;
.PHONY: release/chart-bump/revision

release/chart-bump/minor:
	@for chart in $(EMISSARY_CHART) ; do \
		$(OSS_HOME)/charts/scripts/bump_chart_version.sh minor $$chart/Chart.yaml ; \
	done ;
.PHONY: release/chart-bump/minor

# This is pretty Draconian. Use with care.
chart-clean:
	@for chart in $(EMISSARY_CHART) ; do \
		git restore $$chart/Chart.yaml $$chart/values.yaml && \
			rm -f $$chart/*.tgz $$chart/index.yaml $$chart/tmp.yaml; \
	done ;
.PHONY: chart-clean

$(OSS_HOME)/.circleci/yq:
	cd $(OSS_HOME)/.circleci/yq.d/ && go build -o $(abspath $@) github.com/mikefarah/yq/v3
