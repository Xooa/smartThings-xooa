definition(
    name: "Blockchain Event Viewer",
    namespace: "xooa",
    author: "Arisht Jain",
    description: "Provides information about the state and past events of the specified devices.",
    category: "Convenience",
	iconUrl: "http://cdn.device-icons.smartthings.com/Home/home1-icn.png",
    iconX2Url: "http://cdn.device-icons.smartthings.com/Home/home1-icn@2x.png",
    iconX3Url: "http://cdn.device-icons.smartthings.com/Home/home1-icn@3x.png"){
        appSetting "appId"
        appSetting "apiToken"
    }
 preferences {
 	
	page(name: "indexPage", title: "Enter credentials", nextPage: "mainPage", uninstall: true){
    	section() {
        	input "appId", "text",
				title: "App ID:"
            input "bearer", "text",
                title: "API token:"
        }
    }
    page(name: "mainPage", title: "Your devices", nextPage: "detailPage", install: true, uninstall: true)
    page(name: "detailPage", title: "Past Event Details", install: true, uninstall: true)
}
def mainPage() {
    dynamicPage(name: "mainPage") {
        section() {
        	log.debug "got with settings: ${settings}"
            paragraph "Click on the devices to view full details"
            def appId = settings.appId
            def bearer = settings.bearer
            def json = "{\"args\":["
            json += "]}"
            def params = [
                uri: "https://api.xooa.com/api/${appId}/invoke/getDeviceLastEvent",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "content-type": "application/json",
                    "accept": "application/json"
                ],
                body: json
            ]
            // log.debug(params)
            try {
                httpPostJson(params) { resp ->
                    // log.debug "response data: ${parseJson(resp.data.payload)[0].Record}"
                    for(device in parseJson(resp.data.payload)) {
                    	log.debug(device)
                        def hrefParams = [
                            deviceId: "${device.DeviceId}",
                            name: "${device.Record.displayName}"
                        ]
                        log.debug hrefParams
                        href(name: "toDetailsPage",
                            title: "${device.Record.displayName} - ${device.Record.value}",
                            params: hrefParams,
                            page: "detailPage")
                        // paragraph "${device.Record.displayName}-${device.Record.value}"
                    }
                }
            } catch (groovyx.net.http.HttpResponseException ex) {
                if (ex.statusCode < 200 || ex.statusCode >= 300) {
                    log.debug "Unexpected response error: ${ex.statusCode}"
                    log.debug ex
                    log.debug ex.response.contentType
                }
            }

        }

    }
}

def detailPage(params1) {
    log.debug "params: ${params1}"
    dynamicPage(name: "detailPage") {
        section("${params1?.name}") {
            def appId = settings.appId
            def bearer = settings.bearer
            def json = "{\"args\":[\""
            json += params1?.deviceId
            json += "\"]}"
            def paramaters = [
                uri: "https://api.xooa.com/api/${appId}/invoke/getHistoryForDevice",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "content-type": "application/json",
                    "accept": "application/json"
                ],
                body: json
            ]
            log.debug "did ${params1?.deviceId}"
            if(params1?.deviceId != null) {
            	log.debug paramaters 
                try {
                    httpPostJson(paramaters) { resp ->
                        log.debug "response data: ${resp.data}"
                        if(resp.data.payload){
                            def payload1 = parseJson(resp.data.payload)
                            for(transaction in payload1) {
                                log.debug(transaction)
                                transaction.Record.time = transaction.Record.time.replaceAll('t',' ')
                                log.debug transaction.Record.time
                                def time = transaction.Record.time.take(19)
                                log.debug time
                                paragraph "${time} - ${transaction.Record.value}"
                            }
                        }
                    }
                } catch (groovyx.net.http.HttpResponseException ex) {
                    if (ex.statusCode < 200 || ex.statusCode >= 300) {
                        log.debug "Unexpected response error: ${ex.statusCode}"
                        log.debug ex.response
                        log.debug ex.response.contentType
                    }
                }
            }
        }
    }
}
def installed() {
    log.debug "Installed."

    initialize()
}
def updated() {
    log.debug "Updated."
    initialize()
}
def initialize() {
    log.debug "Initialized"
}
def history() {
    log.debug "history"
    section("Log these presence sensors:") {
        input "presences", "capability.presenceSensor", multiple: true, required: false
    }
}