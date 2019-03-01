 /**
 *  Blockchain Event Viewer
 *
 *  Copyright 2018 Xooa
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at:
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 *  for the specific language governing permissions and limitations under the License.
 */
 /*
 * Original source from Arisht Jain:
 *  https://github.com/xooa/samples
 *
 */

definition(
    name: "Blockchain Event Viewer",
    namespace: "xooa",
    author: "Arisht Jain",
    description: "Provides information about the state and past events of the specified devices.",
    category: "Convenience",
	iconUrl: "https://xooa.com/blockchain/img/logo1.png",
    iconX2Url: "https://xooa.com/blockchain/img/logo1.png",
    iconX3Url: "https://xooa.com/blockchain/img/logo1.png")
    
 preferences {
	page(name: "indexPage", title: "Enter credentials", nextPage: "mainPage", uninstall: true)
    page(name: "mainPage", title: "Your devices", nextPage: "datePage", install: true)
    page(name: "datePage", title: "Select the date", nextPage: "detailPage")
    page(name: "detailPage", title: "Past Event Details", install: true)
}

def indexPage() {
    dynamicPage(name: "indexPage") {
    	section() {
            input "apiToken", "text",
                title: "Xooa Participant API token:"
        }
    }
}

def mainPage() {
    dynamicPage(name: "mainPage") {
        section() {
        	log.debug "settings: ${settings}"
            def apiToken = settings.apiToken

            // queryLocation() function present in smart contract is called in this request. 
            // Modify the endpoint of this URL accordingly if function name is changed
            def params = [
                uri: "https://api.xooa.com/api/v1/xldb/query?selector=%7B%22docType%22%3A%22EventLess%22%7D",
                headers: [
                    "Authorization": "Bearer ${apiToken}",
                    "accept": "text/html",
                    "requestContentType": "text/html",
                    "contentType": "text/html"
                ]
            ]
            log.debug("queryLocation API params: ${params}")
            try {
                httpGet(params) { resp ->
                	log.debug resp.data
                    if(resp.status == 202) {
                        def sleepTime = 3000
                        def requestCount = 5
                        def i = 0
                        def responseStatus = 202
                        while (i < requestCount && responseStatus == 202) {
                            pause(sleepTime)
                            def params1 = [
                                uri: "https://api.xooa.com/api/v1/results/${resp.data.resultId}",
                                headers: [
                                    "Authorization": "Bearer ${bearer}",
                                    "content-type": "application/json"
                                ]
                            ]
                            log.debug("results API params: ${params1}")
                            try {
                                def continueRequest = 0
                                log.debug "Making API request to check for response."
                                httpGet(params1) { resp1 ->
                                    log.debug "response from results API endpoint: ${resp1.data}"
                                    if(resp1.status == 200  && resp1.data.data.size()) {
                                        responseStatus = 200
                                        paragraph "Click on the devices to view full details"
                                        for(device in resp.data.data) {
                                            device.value.time = device.value.time.replaceAll('t',' ')
                                            def time = device.value.time.take(19)
                                            def date = device.value.time.take(10)
                                            def hrefParams = [
                                                deviceId: "${device.key}",
                                                name: "${device.value.displayName}",
                                                date: "${date}"
                                            ]
                                            href(name: "toDatePage",
                                                title: "${device.value.displayName} - ${device.value.value}",
                                                description: "Last updated at: ${time}",
                                                params: hrefParams,
                                                page: "datePage")
                                        }
                                    } else if (resp1.status == 202) {
                                        log.debug "request not processed yet."
                                        i++
                                        continueRequest = 1
                                    }
                                }
                                if(continueRequest == 1){
                                    continue
                                }
                            } catch (groovyx.net.http.HttpResponseException ex) {
                                log.debug "Unexpected response error: ${ex.statusCode}"
                                log.debug ex
                                log.debug ex.response.contentType
                                break
                            }
                        }
                    } else if(resp.data.data.size()){
            			paragraph "Click on the devices to view full details"
                        for(device in resp.data.data) {
                            device.value.time = device.value.time.replaceAll('t',' ')
                            def time = device.value.time.take(19)
                            def date = device.value.time.take(10)
                            def hrefParams = [
                                deviceId: "${device.key}",
                                name: "${device.value.displayName}",
                                date: "${date}"
                            ]
                            href(name: "toDatePage",
                                title: "${device.value.displayName} - ${device.value.value}",
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
                def apiToken = settings.apiToken
                def date = Date.parse("yyyy-MM-dd'T'HH:mm:ss", "${settings.year}-${settings.month}-${settings.day}T00:00:00").format("yyyyMMdd")
                def json = "[\"${state.deviceId}\",\"${date}\"]"

                // queryByDate() function present in smart contract is called in this request. 
                // Modify the endpoint of this URL accordingly if function name is changed
                // Modify the json parameter sent in this request if definition of the function is changed in the smart contract
                def parameters = [
                    uri: "https://api.xooa.com/api/v1/xldb/query?selector=%7B%22docType%22%3A%22Event%22,%22deviceId%22%3A%22${state.deviceId}%22,%22date%22%3A%22${date}%22%7D&fields=value,time",
                    headers: [
                        "Authorization": "Bearer ${apiToken}",
                        "accept": "application/json"
                    ]
                ]
                log.debug parameters
                try {
                    httpGet(parameters) { resp ->
                        log.debug resp.data
                        if(resp.status == 202) {
                            def sleepTime = 3000
                            def requestCount = 5
                            def i = 0
                            def responseStatus = 202
                            while (i < requestCount && responseStatus == 202) {
                                pause(sleepTime)
                                def params1 = [
                                    uri: "https://api.xooa.com/api/v1/results/${resp.data.resultId}",
                                    headers: [
                                        "Authorization": "Bearer ${bearer}",
                                        "content-type": "application/json"
                                    ]
                                ]
                                log.debug("results API params: ${params1}")
                                try {
                                    def continueRequest = 0
                                    log.debug "Making API request to check for response."
                                    httpGet(params1) { resp1 ->
                                        log.debug "response from results API endpoint: ${resp1.data}"
                                        if(resp1.status == 200 && resp1.data.data.size()) {
                                            responseStatus = 200
                                            resp.data.data = resp.data.data.reverse()
                                            for(transaction in resp.data.data) {
                                                transaction.value.time = transaction.value.time.replaceAll('t',' ')
                                                def time = transaction.value.time.take(19)
                                                paragraph "${time} - ${transaction.value.value}"
                                            }
                                        } else if (resp1.status == 202) {
                                            log.debug "request not processed yet."
                                            i++
                                            continueRequest = 1
                                        }
                                    }
                                    if(continueRequest == 1){
                                        continue
                                    }
                                } catch (groovyx.net.http.HttpResponseException ex) {
                                    log.debug "Unexpected response error: ${ex.statusCode}"
                                    log.debug ex
                                    log.debug ex.response.contentType
                                    break
                                }
                            }
                        } else if(resp.data.data.size()){
                            resp.data.data = resp.data.data.reverse()
                            for(transaction in resp.data.data) {
                                transaction.value.time = transaction.value.time.replaceAll('t',' ')
                                def time = transaction.value.time.take(19)
                                paragraph "${time} - ${transaction.value.value}"
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