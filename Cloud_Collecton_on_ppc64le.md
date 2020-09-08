# Using the IBM Cloud Ansible collection on ppc64le

## Install Go Compiler

### Install from Red Hat Software Collections (RHSCL):

Not all RHEL subscriptions include [access to RHSCL][1].

[Hello World - installing Go on RHEL 7][2]

### Installing Go from source

Follow instructuions provided at https://golang.org/doc/install/source.

I chose to [bootstrap from a binary release][3] (1.15.1):

    cd $HOME mkdir tmp wget https://golang.org/dl/go1.15.1.linux-ppc64le.tar.gz
    tar xvzf go1.15.1.linux-ppc64le.tar.gz -C tmp/

[Fetch Git repository][4]:

    sudo yum install git
    git clone https://go.googlesource.com/go goroot
    cd goroot
    git checkout go1.15.1

[Install Go][5]:

    cd src
    GOROOT_BOOTSTRAP=$HOME/tmp/go ./all.bash

Add 'goroot' to your PATH (bash):

    echo "export PATH=$HOME/goroot/bin/:$HOME/go/bin:$PATH" >> $HOME/.bashrc
    source $HOME/.bashrc

## Determine source version requirements

1. Start at the [Ansible Galaxy IBM cloudcollection][6].  Select the desired
   collection version and install via normal process. In the following steps
   set the env var `PROVIDER_VERSION` to the cloudcollection version with a 'v'
   prepended. (e.g.: `PROVIDER_VERSION='v1.11.2'`)

2. Find the [IBM-Cloud/ansible-collection-ibm][7] commit that corresponds to
   the collection version in step 1. Browse the repository at that point in the
   history (the `<>` button) and go to `plugins > module_utils > ibmcloud.py`.
   In the `class Terraform` constructor, `def __init__`, find the default
   `terraform_version` value. [Here's an example][8]. In the following steps
   set the env var `TF_VERSION` to this value with a 'v' prepended. (e.g.:
   `TF_VERSION='v0.12.20'`)

   Note: I'm not sure why git tags aren't utilized in the
   [IBM-Cloud/ansible-collection-ibm][7] repo.


You may also need to create the module tmp directory:

    mkdir -p /var/tmp/ansible/ibmcloud/

## Compile Terraform from source:

    export TF_VERSION='v0.12.20'
    go get github.com/hashicorp/terraform
    cd $HOME/go/src/github.com/hashicorp/terraform
    git checkout $TF_VERSION
    go install github.com/hashicorp/terraform
    terraform -v
    cp $HOME/go/bin/terraform /var/tmp/ansible/ibmcloud/terraform

## Compile IBM Cloud Terraform provider from source:

    export PROVIDER_VERSION='v1.11.2'
    go get github.com/IBM-Cloud/terraform-provider-ibm
    cd $HOME/go/src/github.com/IBM-Cloud/terraform-provider-ibm
    git checkout $PROVIDER_VERSION
    go install github.com/IBM-Cloud/terraform-provider-ibm
    cp $HOME/go/bin/terraform-provider-ibm /var/tmp/ansible/ibmcloud/terraform-provider-ibm_$PROVIDER_VERSION

[1]: https://access.redhat.com/solutions/472793
[2]: https://developers.redhat.com/HW/Go-RHEL-7#troubleshooting_and_faq
[3]: https://golang.org/doc/install/source#bootstrapFromBinaryRelease
[4]: https://golang.org/doc/install/source#fetch
[5]: https://golang.org/doc/install/source#install
[6]: https://galaxy.ansible.com/ibm/cloudcollection
[7]: https://github.com/IBM-Cloud/ansible-collection-ibm/commits/master
[8]: https://github.com/IBM-Cloud/ansible-collection-ibm/blob/c706e66310a154c0aa7b3ecc870ba455a2cf0975/plugins/module_utils/ibmcloud.py#L635
[9]: https://github.com/IBM-Cloud/terraform-provider-ibm/releases
