.PHONY: git
push:
	git config --global url."git://github.com/".insteadOf "https://github.com/"
	git add .
	git commit -m "$m"
	git push

pull:
	git pull
