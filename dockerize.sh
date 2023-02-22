docker build -f Dockerfile -t walenpiotr/url-shortener .
docker tag walenpiotr/url-shortener walenpiotr/url-shortener:1.0.0
docker push walenpiotr/url-shortener:1.0.0