%global debug_package %{nil}
%global __strip /bin/true

Name:          nier 	
Version:        %{ver}
Release:        %{rel}%{?dist}

Summary:	nier is a monitor software,used to monit ceph status and host status.

Group:		SDS
License:	GPL
URL:		http://github.com/journeymidnight
Source0:	%{name}-%{version}-%{rel}.tar.gz
BuildRoot:	%(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)
#BuildRequires:  
Requires:       librados2-devel
Requires:       librados2

%description


%prep
%setup -q -n %{name}-%{version}-%{rel}


%build
#The go build still use source code in GOPATH/src/legitlab/yig/
#keep git source tree clean, better ways to build?
#I do not know
make

%install
rm -rf %{buildroot}
install -D -m 755 %{_builddir}/nier-%{version}-%{rel}/nier %{buildroot}%{_bindir}/nier
install -D -m 644 package/nier.service   %{buildroot}/usr/lib/systemd/system/nier.service
install -D -m 644 conf/basic_model.conf   %{buildroot}/etc/nier/basic_model.conf
install -D -m 644 conf/conf.sample.toml   %{buildroot}/etc/nier/conf.toml

#ceph confs ?

%post
systemctl enable nier
systemctl restart nier


%preun

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root,-)
/usr/lib/systemd/system/nier.service
/usr/bin/nier
/etc/nier/basic_model.conf
/etc/nier/conf.toml

%changelog
