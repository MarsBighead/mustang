{{group_names}}
{% for k in group_names %}
{% if k.startswith('semi') %}
{% for h in groups['semi1'] %}
{% if hostvars[h]['isPrimary'] is defined %}
change master to master_host='{{h}}', master_user='repl', master_password='password', master_port=3306, master_auto_position=1;
{% endif %}
{% endfor %}
{% endif %}
{% endfor %}
