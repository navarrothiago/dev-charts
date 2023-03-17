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


