rsync -r ./ ../cld/ \
--exclude dev \
--exclude departmental \
--exclude docker \
--exclude db \
--exclude .gitignore \
--exclude .git \
--progress --dry-run
