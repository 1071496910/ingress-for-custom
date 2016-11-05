# Usage

## ingress.yaml

```
spec:
  rules:
  - host: hybird.meizu.com
    http:
      paths:
      - backend:
          serviceName: venice-api-v1-bc69a
          servicePort: 9100
        path: /pubapi
  - host: hpc.meizu.com-custom          -custom表示自定义    下面是自定义的方式
    http:
      paths:
      - backend:
          serviceName: fe1-v1-03c1b     后端service 配置不变
          servicePort: 8082
        path: /default/custom           path 指向configmap的/namespace/name    
      - backend:
          serviceName: test-only-bc69a  
          servicePort: 3010
        path: /default/custom
```

## customConfig.yaml

```
data:
  custom: |
    server {
        server_name sms.admin.meizu.com;
       listen 80;
       #access_log  /data/log/nginx/sms_access.log main;


       location ~* .*\.svn.* {
       return 404;
       }

       location / {
           proxy_pass http://silent-update-fe1-v1-03c1b-8082;
                    ！！！格式namespace-service-clusterport
       }



       location /resources {
           proxy_pass http://silent-update-fe1-v1-03c1b-8082;
       }


       location /service/admin {

           proxy_store off;
           proxy_redirect off;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header Host $http_host;
           proxy_pass http://silent-update-fe1-v1-03c1b-8082;
           #proxy_pass http://192.168.8.166:8080/service/admin;

       }

    }
```
