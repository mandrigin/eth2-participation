#/bin/sh

# ssh -NT -L 127.0.0.1:4000:localhost:4000 devops@178.128.155.68

ssh -i ~/.ssh/cloudsigma.key -NT -L 127.0.0.1:4000:10.46.118.157:3500 cloudsigma@31.171.250.161
