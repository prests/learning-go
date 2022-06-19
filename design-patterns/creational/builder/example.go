package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()

	// TODO: Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(10)
	bldr.SetMessage("AYOOOOO")
	bldr.SetType("alert")

	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		panic("Error creating notification")
	}

	fmt.Printf("%v\n", notif)
}
