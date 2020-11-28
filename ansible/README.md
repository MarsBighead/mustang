# Ansible task example

## Skip execute

Set the host configure with `-i`
 
```bash
ansible-playbook -i /etc/ansible/hosts  /srv/yaml/appoint-center.yml --step --start-at-task='start jacoco'
```

## Execute task

Deploy command line

```sh
ansible-playbook site.yml -i hosts
```
