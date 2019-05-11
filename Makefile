.PHONY: git
push:
	git add .
	git commit -m "$m"
	git push

pull:
	git pull
