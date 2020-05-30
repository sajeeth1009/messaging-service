echo "make sure the config file is at the define path"
docker run --env-file email-client-service-env.list -p 5005:5005 github.com/influenzanet/email-client-service:$1