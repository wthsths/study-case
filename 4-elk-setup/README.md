###
- helm upgrade --install --create-namespace elasticsearch elastic/elasticsearch -f es-values.yaml -n logging
- helm upgrade --install --create-namespace kibana elastic/kibana -f kibana-values.yaml -n logging
- helm upgrade --install --create-namespace logstash elastic/logstash -f logstash-values.yaml -n logging
- helm upgrade --install --create-namespace filebeat elastic/filebeat -f filebeat-values.yaml -n logging
