# Docker compose

Esta pasta contém os manifestos para o docker-compose. O docker-compose é uma ferramenta para definir e executar aplicativos Docker com vários contêineres. O arquivo `docker-compose.yml` contém os serviços que serão executados.

## Serviços

- Nginx (Proxy reverso)
- Redmine (Issue Tracker)
- Gitlab (Git)
- :construction: Nextcloud (File Sharing)

## Como usar

O docker-compose deve ser exeuctado na pasta `docker` para que o arquivo `.env` seja carregado. O arquivo `.env` contém as variáveis de ambiente que são usadas nos manifestos.

Commandos úteis:
- `docker-compose up -d` para iniciar os serviços;
- `docker-compose down` para parar os serviços;
- `docker-compose logs -f` para acompanhar os logs dos serviços;

## Plugins não instalados

- redmine_checkbox_wiki
- redmine_checklists
- redmine_dmsf
- redmine_embedded_video
- redmine_latex_mathjax
- redmine_monitoring_controlling
- redmine_user_mentions
- redmine_wiki_gchart_formula

## TODO

- [ ] Recuperar os plugins do redmine, principalmente o redmine_dmsf;
