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
	page(name: "indexPage", title: "Enter credentials", nextPage: "mainPage", uninstall: true)
    page(name: "mainPage", title: "Your devices", nextPage: "datePage", install: true)
    page(name: "datePage", title: "Select the date", nextPage: "detailPage")
    page(name: "detailPage", title: "Past Event Details", install: true)
}

def indexPage() {
    dynamicPage(name: "indexPage") {
    	app.updateSetting("Lid", location.id)
    	section() {
            input "appId", "text",
                title: "Xooa app ID:", submitOnChange: true
            input "apiToken", "text",
                title: "Xooa Participant API token:", submitOnChange: true
            input "locationid", "text",
                title: "Location ID:", submitOnChange: true, defaultValue: location.id
            paragraph "You can share your location id with another user to share events of your devices."
            input "Lid", "text",
                title: "Your Location ID:", defaultValue: location.id
        }
    }
}

def mainPage() {
    dynamicPage(name: "mainPage") {
        section() {
        	log.debug "settings: ${settings}"
            def appId = settings.appId
            def apiToken = settings.apiToken
            // queryLocation() function present in chaincode is called in this request. 
            // Modify the endpoint of this URL accordingly if function name is changed
            def params = [
                uri: "https://api.xooa.com/api/${appId}/query/queryLocation?args=%5B%22${settings.locationid}%22%5D",
                headers: [
                    "Authorization": "Bearer ${apiToken}",
                    "accept": "text/html",
                    "requestContentType": "text/html",
                    "contentType": "text/html"
                ]
            ]
            try {
                httpGet(params) { resp ->
                	log.debug resp.data
                	if(resp.data.size()){
            			paragraph "Click on the devices to view full details"
                        for(device in resp.data) {
                            device.Record.time = device.Record.time.replaceAll('t',' ')
                            def time = device.Record.time.take(19)
                            def date = device.Record.time.take(10)
                            def hrefParams = [
                                deviceId: "${device.Key}",
                                name: "${device.Record.displayName}",
                                date: "${date}"
                            ]
                            href(name: "toDatePage",
                                title: "${device.Record.displayName} - ${device.Record.value}",
                                description: "Last updated at: ${time}",
                                params: hrefParams,
                                page: "datePage")
                        }
                 	} else {
                    	paragraph "No devices found."
                    }
                }
            } catch (groovyx.net.http.HttpResponseException ex) {
               	if (ex.statusCode < 200 || ex.statusCode >= 300) {
                    log.debug "Unexpected response error: ${ex.statusCode}"
                    log.debug ex
                    log.debug ex.response.data
                    log.debug ex.response.contentType
                }
            }

        }

    }
}

def datePage(params1) {
	log.debug "params1: ${params1}"
    dynamicPage(name: "datePage") {
        section() {
            if(params1?.date != null) {
            	def date = params1?.date
                state.deviceName = params1?.name
                state.deviceId = params1?.deviceId
            	date = date.split("-")
                app.updateSetting("day", date[2])
                app.updateSetting("month", date[1])
                app.updateSetting("year", date[0])
                input name: "day", type: "number", title: "Day", required: true
                input name: "month", type: "number", title: "Month", required: true
                input name: "year", type: "number", description: "Format(yyyy)", title: "Year", required: true
            } 
            else {
                input name: "day", type: "number", title: "Day", required: true
                input name: "month", type: "number", title: "Month", required: true
                input name: "year", type: "number", description: "Format(yyyy)", title: "Year", required: true
            }
        }
    }
}

def detailPage() {
    dynamicPage(name: "detailPage") {
        section("${state.deviceName}") {
            log.debug "did: ${state.deviceId}"
            if(state.deviceId != null) {
                def appId = settings.appId
                def apiToken = settings.apiToken
                def date = Date.parse("yyyy-MM-dd'T'HH:mm:ss", "${settings.year}-${settings.month}-${settings.day}T00:00:00").format("yyyyMMdd")
                def json = "%5B%22${settings.locationid}%22,%22${state.deviceId}%22,%22${date}%22%5D"
                // queryByDate() function present in chaincode is called in this request. 
                // Modify the endpoint of this URL accordingly if function name is changed
                // Modify the json parameter sent in this request if definition of the function is changed in the chaincode
                def paramaters = [
                    uri: "https://api.xooa.com/api/${appId}/query/queryByDate?args=${json}",
                    headers: [
                        "Authorization": "Bearer ${apiToken}",
                        "accept": "application/json"
                    ]
                ]
                try {
                    httpGet(paramaters) { resp ->
                        log.debug resp.data
                        if(resp.data.size()){
                            resp.data = resp.data.reverse()
                            for(transaction in resp.data) {
                                transaction.Record.time = transaction.Record.time.replaceAll('t',' ')
                                def time = transaction.Record.time.take(19)
                                paragraph "${time} - ${transaction.Record.value}"
                            }
                        } else {
                            paragraph "No events found for the selected date."
                        }
                    }
                } catch (groovyx.net.http.HttpResponseException ex) {
                    if (ex.statusCode < 200 || ex.statusCode >= 300) {
                        log.debug "Unexpected response error: ${ex.statusCode}"
                        log.debug ex.response
                        log.debug ex.response.contentType
                    }
                }
           	} else {
            	paragraph "Unable to retrieve device ID"
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