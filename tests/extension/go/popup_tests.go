package main

import (
	"fmt"
	"github.com/Archs/chrome/api/alarms"
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/chrome/api/browserAction"
	"github.com/Archs/chrome/api/browsingData"
	"github.com/Archs/chrome/api/cookies"
	"github.com/Archs/chrome/api/extension"
	"github.com/Archs/chrome/api/fontSettings"
	"github.com/Archs/chrome/api/history"
	"github.com/Archs/chrome/api/tabs"
	"time"

	QUnit "github.com/fabioberger/qunit"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var (
	doc = dom.GetWindow().Document()
)

func main() {

	QUnit.Module("Chrome-Popup")

	/*
	* Alarm Method Tests
	 */

	// Create New Alarm
	alarmOps := js.M{
		"when": time.Now().UnixNano() + 1000000,
	}
	alarms.Create("test_alarm", alarmOps)

	// Get the Alarm created above
	alarms.Get("test_alarm", func(alarm alarms.Alarm) {
		QUnit.Test("Alarm.Get()", func(assert QUnit.QUnitAssert) {
			assert.Equal(alarm.Name, "test_alarm", "Get")
		})

		// Clear the Alarm retrieved above
		alarms.Clear("test_alarm", func(wasCleared bool) {
			QUnit.Test("Alarm.Clear()", func(assert QUnit.QUnitAssert) {
				assert.Equal(wasCleared, true, "Clear")
			})

			// Create two more Alarms
			alarms.Create("test_alarm2", alarmOps)
			alarms.Create("test_alarm3", alarmOps)

			// Get both Alarms created above
			alarms.GetAll(func(alms []alarms.Alarm) {
				QUnit.Test("Alarm.GetAll()", func(assert QUnit.QUnitAssert) {
					assert.Equal(alms[0].Name, "test_alarm2", "GetAll")
					assert.Equal(alms[1].Name, "test_alarm3", "GetAll")
				})
				// Clear both Alarms above
				alarms.ClearAll(func(wasCleared bool) {
					QUnit.Test("Alarm.ClearAll()", func(assert QUnit.QUnitAssert) {
						assert.Equal(wasCleared, true, "Clear")
					})
				})
			})
		})
	})

	/*
	* Bookmarks Method Test
	 */
	bookmark := js.M{
		"title": "Testing",
		"url":   "http://www.testing.com/",
	}
	// Test Create New Bookmarks
	bookmarks.Create(bookmark, func(result bookmarks.TreeNode) {
		QUnit.Test("Bookmarks.Create()", func(assert QUnit.QUnitAssert) {
			assert.Equal(result.Title, "Testing", "Create")
		})

		// Test Get Bookmark by List of Id's
		bookmarks.Get([]string{result.Id}, func(results []bookmarks.TreeNode) {
			QUnit.Test("Bookmarks.Get()", func(assert QUnit.QUnitAssert) {
				assert.Equal(results[0].Url, "http://www.testing.com/", "Get")
			})
		})
	})

	/*
	* BrowserAction Method Tests
	 */

	change := js.M{
		"title": "Testing",
	}
	browserAction.SetTitle(change)
	browserAction.GetTitle(js.M{}, func(result string) {
		QUnit.Test("BrowserAction.GetTitle()", func(assert QUnit.QUnitAssert) {
			assert.Equal(result, "Testing", "GetTitle")
		})
	})

	/*
	* BrowsingData Method Tests
	 */

	// Test Retrieving BrowserData Settings
	browsingData.Settings(func(result js.M) {
		for key, _ := range result {
			QUnit.Test("BrowsingData.Settings()", func(assert QUnit.QUnitAssert) {
				assert.Equal(key, "dataRemovalPermitted", "Settings")
			})
			break
		}
	})

	/*
	* Cookies Method Tests
	 */

	// Set a new Cookie
	cookieInfo := js.M{
		"url":   "http://www.google.com",
		"name":  "testing",
		"value": "testvalue",
	}
	cookies.Set(cookieInfo, func(cookie cookies.Cookie) {
		QUnit.Test("Cookies.Set()", func(assert QUnit.QUnitAssert) {
			assert.Equal(cookie.Name, "testing", "Set")
			assert.Equal(cookie.Value, "testvalue", "Set")
		})

		// Get the cookie set previously
		cookieInfo = js.M{
			"url":  "http://www.google.com",
			"name": "testing",
		}
		cookies.Get(cookieInfo, func(cookie cookies.Cookie) {
			QUnit.Test("Cookies.Get()", func(assert QUnit.QUnitAssert) {
				assert.Equal(cookie.Name, "testing", "Get")
			})
		})
	})

	/*
	* Extension Method Tests
	 */

	// Get the popup views for the Extension
	fetchProperties := js.M{
		"type": "popup",
	}
	windows := extension.GetViews(fetchProperties)
	QUnit.Test("Extension.GetViews()", func(assert QUnit.QUnitAssert) {
		assert.Equal(windows[0].Incognito, false, "GetViews")
	})

	/*
	* FontSettings Method Tests
	 */

	// Get Details for a Generic Font Family
	fontDetails := js.M{
		"genericFamily": "standard",
		"script":        "Arab",
	}
	fontSettings.GetFont(fontDetails, func(details js.M) {
		for key, _ := range details {
			QUnit.Test("FontSettings.GetFont()", func(assert QUnit.QUnitAssert) {
				assert.Equal(key, "fontId", "GetFont")
			})
			break
		}
	})

	/*
	* History Method Tests
	 */

	// Add URL to History
	urlDetails := js.M{
		"url": "http://www.testing.com/",
	}
	history.AddUrl(urlDetails, func() {

		// Search for created URL in History
		s := js.M{
			"text": "www.testing.com",
		}
		history.Search(s, func(results []history.HistoryItem) {
			QUnit.Test("History.Search()", func(assert QUnit.QUnitAssert) {
				assert.Equal(results[0].Url, "http://www.testing.com/", "Search")
			})
		})
	})

	/*
	* Tab Method Test
	 */

	// Find the current Tab
	queryInfo := js.M{
		"currentWindow": true,
		"active":        true,
	}
	tabs.Query(queryInfo, func(tbs []tabs.Tab) {
		id := tbs[0].Id

		// Send Message to the given Tab & have Event Listener on Tab Respond
		msg := js.M{"greeting": "hello"}
		tabs.SendMessage(id, msg, func(response js.M) {
			err := js.Global.Get("chrome").Get("runtime").Get("lastError")
			if err.String() != "undefined" {
				fmt.Println("Tabs.SendMessage Error: ", err.Get("message").String())
			}
			QUnit.Test("Tabs.SendMessage() & Runtime.OnMessage() Event", func(assert QUnit.QUnitAssert) {
				assert.Equal(response["farewell"], "goodbye", "SendMessage")
			})
		})
	})

}
