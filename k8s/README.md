# Getting Started

Esta pasta contém os arquivos de configuração para o Kubernetes. O Kubernetes é um sistema de orquestração de contêineres que permite a implantação e gerenciamento de aplicativos em contêineres em um cluster de computadores.

:construction: Este projeto está em construção. :construction:

A ideia é permitir que os serviços sejam implantados em um cluster de Kubernetes. São eles:

- Longhorn (Storage)
- Redmine (Issue Tracker)
- Nextcloud (File Sharing)

Pré-requisitos:

- [k3s](https://k3s.io/)
- [helm](https://helm.sh/)

A pasta `images` contenha script para importar as imagens necessárias para o Kubernetes, usado em nodes que não possuem acesso à internet (air-gapped).

## Deployments

### Longhorn

```bash
helm install longhorn longhorn/ -n longhorn-system --create-namespace
```

### Redmine

```bash
helm install dev-redmine dev/ \
--set redmine.enabled=true \
--set kube-prometheus-stack.enabled=false \
--set redmine.smtpHost="10.13.148.1" \
--set redmine.smtpPort=25 \
--set redmine.smtpUser="ctex_rds@3cta.eb.mil.br" \
--set redmine.smtpPassword=<smtp_pass> \
--set redmine.image.repository=redmine \
--set redmine.image.tag=latest
```


### :construction: Nextcloud and Gitlab


