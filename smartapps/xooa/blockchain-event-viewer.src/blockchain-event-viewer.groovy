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
            def bearer = settings.bearerdefinition(
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
            def params = [
                uri: "https://api.xooa.com/api/${appId}/query/getDeviceLastEvent?args=%5B%5D",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "accept": "application/json"
                ]
            ]
            try {
                httpGet(params) { resp ->
                    for(device in resp.data) {
                        def hrefParams = [
                            deviceId: "${device.DeviceId}",
                            name: "${device.Record.displayName}"
                        ]
                        device.Record.time = device.Record.time.replaceAll('t',' ')
                        def time = device.Record.time.take(19)
                        href(name: "toDetailsPage",
                            title: "${device.Record.displayName} - ${device.Record.value}",
                            description: "Last updated at: ${time}",
                            params: hrefParams,
                            page: "detailPage")
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
            def json = "%5B%22${params1?.deviceId}%22%5D"
            def paramaters = [
                uri: "https://api.xooa.com/api/${appId}/query/getHistoryForDevice?args=${json}",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "accept": "application/json"
                ]
            ]
            log.debug "did: ${params1?.deviceId}"
            if(params1?.deviceId != null) {
                try {
                    httpGet(paramaters) { resp ->
                        if(resp.data){
                            resp.data = resp.data.reverse()
                            for(transaction in resp.data) {
                                transaction.Record.time = transaction.Record.time.replaceAll('t',' ')
                                def time = transaction.Record.time.take(19)
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
            def params = [
                uri: "https://api.xooa.com/api/${appId}/query/getDeviceLastEvent?args=%5B%5D",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "accept": "application/json"
                ]
            ]
            try {
                httpGet(params) { resp ->
                    for(device in resp.data) {
                        def hrefParams = [
                            deviceId: "${device.DeviceId}",
                            name: "${device.Record.displayName}"
                        ]
                        device.Record.time = device.Record.time.replaceAll('t',' ')
                        def time = device.Record.time.take(19)
                        href(name: "toDetailsPage",
                            title: "${device.Record.displayName} - ${device.Record.value}",
                            description: time,
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
            def json = "%5B%22${params1?.deviceId}%22%5D"
            def paramaters = [
                uri: "https://api.xooa.com/api/${appId}/query/getHistoryForDevice?args=${json}",
                headers: [
                    "Authorization": "Bearer ${bearer}",
                    "accept": "application/json"
                ]
            ]
            log.debug "did: ${params1?.deviceId}"
            if(params1?.deviceId != null) {
                try {
                    httpGet(paramaters) { resp ->
                        if(resp.data){
                            resp.data = resp.data.reverse()
                            for(transaction in resp.data) {
                                transaction.Record.time = transaction.Record.time.replaceAll('t',' ')
                                def time = transaction.Record.time.take(19)
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