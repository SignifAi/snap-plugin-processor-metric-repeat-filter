Name:           signifai-snap-plugin-processor-metric-repeat-filter
Version:        VERSION
Release:        1%{?dist}
Summary:        Snap Processor Plugin to Filter Repeat Metrics

License:        Apache
URL:            https://github.com/signifai/snap-plugin-processor-metric-repeat-filter
Source0:        snap-plugin-processor-metric-repeat-filter

Requires:       signifai-go >= 1.8.3-el6.1
Requires:       signifai-snap-agent >= 1.2.0-el6.1

%description


%prep
# No prep; already done

%build
# No build; we already did that.

%install
rm -rf $RPM_BUILD_ROOT

mkdir -p $RPM_BUILD_ROOT/opt/signifai/snap/plugins
cp %{SOURCE0} $RPM_BUILD_ROOT/opt/signifai/snap/plugins/snap-plugin-processor-metric-repeat-filter
%clean


%files
%defattr(-,root,root,-)
/opt/signifai/snap/plugins/snap-plugin-processor-metric-repeat-filter

%changelog
