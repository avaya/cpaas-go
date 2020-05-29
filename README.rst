cpaas-go
==========

This golang package is an open source tool built to simplify interaction with
the `Avaya CPaaS <http://www.avayacloud.com>`_ telephony platform. Avaya OneCloud™️ CPaaS  makes adding voice
and SMS to applications fun and easy.

For more information about Avaya CPaaS, please visit:
`Avaya OneCloud™️ CPaaS  <https://www.https://www.avaya.com/en/products/cloud/cpaas/>`_

To read the official documentation, please visit: `Avaya CPaaS Docs <http://docs.avayacloud.com/aspx/docs>`_.


Installation
============

Clone the repo, and install via ``go get``:

.. code-block:: bash

    $ go get -u github.com/avaya/cpaas-go
    
.. code-block:: bash

Usage
======

Authorization
----
In order to authorize against Avaya OneCloud™️ CPaaS  services you'll have to define authentication credentials.

.. code-block:: bash

    export ZANG_CLOUD_ACCOUNT_SID="{YourAccountSid}"
    export ZANG_CLOUD_AUTH_TOKEN="{YourAccessToken}"
    
.. code-block:: bash

Import
----

.. code-block:: golang
    
    import (
        zang "github.com/avaya/cpaas-go"
    )

.. code-block:: golang

Logging
----

In order to use the Avaya OneCloud CPaaS library, logging must be initialized:

.. code-block:: golang

    package main
    import (
    	log "github.com/sirupsen/logrus"
    	zang "github.com/avaya/cpaas-go"
    )
    func init() {
    	log.SetLevel(log.DebugLevel)
    	formatter := &log.TextFormatter{
    		FullTimestamp: true,
    	}
	    log.SetFormatter(formatter)
    }
    
.. code-block:: golang

REST
----

See the `Avaya CPaaS REST API documentation <http://docs.avayacloud.com/aspx/rest>`_
for more information.

**NOTE: ** Please go through tests for specific endpoint to see the example

Send SMS Example
----------------

.. code-block:: golang

    func sendsms() {
    	client, err := zang.NewClient()
    	if err != nil {
	    	log.Errorln("Client creation failed:", err)
	    	return
	    }
	    log.Debugln("Sending request")
	    response, err := client.SendSms(map[string]string{
    		"From": "E164 From",
    		"To":   "E164 To",
    		"Body": "Test Body,
    	})
        
     func main() {
    	os.Setenv("ZANG_CLOUD_ACCOUNT_SID", "{YourAccountSid}")
    	os.Setenv("ZANG_CLOUD_AUTH_TOKEN", "{YourAccessToken}")
	    sendsms()
    }

.. code-block:: golang

InboundXML
==========

InboundXML is an XML dialect which enables you to control phone call flow.
For more information please visit the `Avaya CPaaS InboundXML documentation
<http://docs.avayacloud.com/aspx/inboundxml>`_.

<Say> Example
-------------

.. code-block:: golang

  ixml, err := New(Response{Say: &Say{
    Voice: "female",
    Value: "Welcome to Avaya CPaaS!",
    Loop:  3,
  }})

  fmt.Print(ixml)

.. code-block:: golang

will render

.. code-block:: xml

    <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <Response>
        <Say loop="3" voice="female" language="en">Welcome to Avaya CPaaS!</Say>
    </Response>

.. code-block:: xml
