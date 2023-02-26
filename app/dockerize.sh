docker build -f Dockerfile.prod -t walenpiotr/url-shortener .
docker tag walenpiotr/url-shortener walenpiotr/url-shortener:1.1.3
docker push walenpiotr/url-shortener:1.1.3