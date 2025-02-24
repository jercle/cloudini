for FILE in *.json; do
  awk '{ gsub("oldstr","newstr"); print }' ${FILE} > ./output/${FILE}
done
