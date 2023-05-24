package HotRemote

const (
	ListDevicesAction = "GetMultiRoom"
)

func ButtonPressRequest(TargetBox string, press int) Request {
	return Request{
		Action:    "ButtonEvent",
		Press:     press,
		TargetBox: TargetBox,
	}
}

func GotoLiveRequest(TargetBox string, press string, params string) Request {
	return Request{
		Action:    "CustomEvent",
		Press:     press,
		TargetBox: TargetBox,
		Event:     "GotoLive",
		Params:    params,
	}
}

func ListDevicesRequest() Request {
	return Request{
		Action:    ListDevicesAction,
		Press:     "getDevices",
		TargetBox: "SERVER",
	}
}

func (c *Client) ListDevices() ([]Device, error) {
	response := &ListDevicesResponse{}

	if err := c.send(ListDevicesRequest(), response); err != nil {
		return nil, err
	}

	var devices []Device
	for _, item := range response.Data {
		devices = append(devices, Device{
			ID:   item[0].(string),
			Name: item[1].(string),
		})
	}

	return devices, nil
}

func (c *Client) ButtonPress(targetBox string, buttonID int) error {
	response := Request{}
	// the server just echos the request we sent.
	return c.send(ButtonPressRequest(targetBox, buttonID), &response)
}

func (c *Client) CustomEvent(targetBox string, buttonID string, params string) error {
	response := Request{}
	// the server just echos the request we sent.
	return c.send(GotoLiveRequest(targetBox, buttonID, params), &response)
}
