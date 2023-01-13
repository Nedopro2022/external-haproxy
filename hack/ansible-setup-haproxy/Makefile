.PHONY: requirements
requirements:
	ansible-galaxy collection install -r requirements.yml

.PHONY: setup
setup:
	ansible-playbook -i inventory --user=ubuntu --ask-become-pass 1-setup.yml

.PHONY: config
config:
	ansible-playbook -i inventory --user=ubuntu --ask-become-pass 2-config.yml
