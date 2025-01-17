package extension

const (
	tplExtensionYaml = `apiVersion: v1
name: [[ .Name ]]
version: 0.1.0
displayName:
  zh: 示例扩展组件
  en: Sample Extension
description:
  zh: 这是一个示例扩展组件，这是它的描述
  en: This is a sample extension, and this is its description
keywords:
  - [[ .Category ]]
home: https://kubesphere.io
sources:
  - https://github.com/kubesphere
kubeVersion: ">=1.19.0"
ksVersion: ">=3.0.0"
dependencies:
  - name: frontend
    condition: frontend.enabled
  - name: backend
    condition: backend.enabled
icon: ./favicon.svg
`
	tplPermissionsYaml = `kind: ClusterRole
rules:
  - verbs:
      - 'create'
      - 'patch'
      - 'update'
      - 'get'
    apiGroups:
      - 'extensions.kubesphere.io'
    resources:
      - '*'

---
kind: Role
rules:
  - verbs:
      - '*'
    apiGroups:
      - ''
      - 'apps'
      - 'batch'
      - 'app.k8s.io'
      - 'autoscaling'
    resources:
      - '*'
  - verbs:
      - '*'
    apiGroups:
      - 'networking.k8s.io'
    resources:
      - 'ingresses'
      - 'networkpolicies'
`
	tplValuesYaml = `frontend:
  enabled: true
  image:
    repository: kubespheredev/employee-frontend
    tag: latest

backend:
  enabled: true
  image:
    repository: kubespheredev/employee-api
    tag: latest
`
	tplHelmignore = `# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*.orig
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`
	tplFavicon = `<svg xmlns="http://www.w3.org/2000/svg" width="477.333" height="477.333"><svg version="1.0" xmlns="http://www.w3.org/2000/svg" width="477.333" height="477.333" viewBox="0 0 358 358" fill="#00ad6e"><path d="M174.5 5.4c-1.1.8-28.5 16.8-61 35.6-32.4 18.9-65.2 37.9-72.8 42.4l-13.7 8v175.3l12.8 7.4c7 4.1 13.6 7.9 14.7 8.5 1.1.6 15.3 8.9 31.5 18.4 16.2 9.4 31.3 18.2 33.5 19.5 2.2 1.2 15.9 9.2 30.5 17.8 14.6 8.5 26.6 15.4 26.8 15.2.2-.2.4-114.8.2-129 0-3.8-.3-4.1-6.2-7.5-3.5-1.9-8.3-4.7-10.8-6.2s-5.1-3.1-5.9-3.4c-.8-.4-8.6-5-17.5-10.2-20.1-11.8-25.5-15-28.3-16.4-1.3-.6-2.3-1.4-2.3-1.7 0-.5 7-4.8 22.5-13.6 7-4 31.5-18.3 37.4-21.8 2.5-1.5 6.1-3.5 7.9-4.4 2.3-1.2 3.2-2.4 3.2-4.2.4-21.1.4-131.1 0-131-.3 0-1.4.6-2.5 1.3zm40 64.8c0 29.4.4 45.3 1 45.3.9 0 14.9-7.6 19.5-10.5 1.4-.9 8.6-5.1 15.7-9.1 14-7.9 41.9-24.2 42.5-24.7.6-.6-14.5-9.6-42.1-25.3-2.5-1.5-11.7-6.7-20.3-11.8-8.6-5-15.8-9.1-16-9.1-.2 0-.3 20.4-.3 45.2z"/><path d="M295.5 111.5c-18.7 10.7-34.7 19.8-35.5 20.3-.8.5-8.7 5-17.5 9.9-8.8 5-26.1 14.8-38.5 21.9-12.4 7-23.5 13.3-24.8 13.9-2.9 1.6-3.7 1 14.8 11.2 1.9 1.1 15.1 8.6 29.3 16.7 14.1 8 27.4 15.6 29.5 16.7 2 1.2 10.5 6 18.7 10.7 8.3 4.7 17 9.6 19.5 11 2.5 1.4 12.2 6.9 21.7 12.3 9.5 5.5 17.6 9.9 18 9.9 1.2 0 1-173.3-.2-173.6-.6-.1-16.3 8.5-35 19.1zm-81 176c0 25 .1 45.5.2 45.5.3 0 3.7-2 16.8-9.5 23-13.3 29.1-16.8 31.5-18.4 1.4-.8 2.7-1.6 3-1.6.3 0 1.6-.8 3-1.7 1.4-.8 7.5-4.4 13.5-7.8 6.1-3.5 11.1-6.4 11.3-6.5.4-.3-77.5-45.5-78.5-45.5-.4 0-.8 20.5-.8 45.5z"/></svg><style>@media (prefers-color-scheme:light){:root{filter:none}}</style></svg>`

	tplReadmeZh = `这是一个示例扩展组件，这里展示了他的详细介绍，你可以在这里使用 Markdown 语法编写内容。`
	tplReadmeEn = `This is a sample extension, which is shown in more detail here, and you can write it here using Markdown syntax.`
)
