install:
	git clone --depth 1 --branch v2.26.0 https://github.com/kubernetes-sigs/kubespray.git kubespray
	cd kubespray && \
	    python3 -m venv venv && \
		. venv/bin/activate && pip3 install -r requirements.txt

copy:
	rm -rf kubespray/inventory/mycluster
	cp -r inventory kubespray/inventory/mycluster

clean:
	rm -rf kubespray
