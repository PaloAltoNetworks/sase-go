package routing

/*
Config ""

Parent chains:
* "Config"

Args:

Param AcceptRouteOverSC (bool): The AcceptRouteOverSC param.

Param AddHostRouteToIkePeer (bool): The AddHostRouteToIkePeer param.

Param BackboneRouting (string): The BackboneRouting param. String values: []string{"no-asymmetric-routing", "asymmetric-routing-only", "asymmetric-routing-with-load-share"}

Param OutboundRoutesForServices (list of strings): The OutboundRoutesForServices param.

Param RoutingPreference (RoutingPreferenceObject): The RoutingPreference param. RoutingPreferenceObject instance.

Param WithdrawStaticRoute (bool): The WithdrawStaticRoute param.
*/
type Config struct {
    AcceptRouteOverSC bool `json:"accept_route_over_SC,omitempty"`
    AddHostRouteToIkePeer bool `json:"add_host_route_to_ike_peer,omitempty"`
    BackboneRouting string `json:"backbone_routing,omitempty"`
    OutboundRoutesForServices []string `json:"outbound_routes_for_services,omitempty"`
    RoutingPreference *RoutingPreferenceObject `json:"routing_preference,omitempty"`
    WithdrawStaticRoute bool `json:"withdraw_static_route,omitempty"`
}

/*
RoutingPreferenceObject ""

Parent chains:
* "Config routing_preference"

Args:

Param Default (interface{}): The Default param. interface{} instance.

Param HotPotatoRouting (interface{}): The HotPotatoRouting param. interface{} instance.
*/
type RoutingPreferenceObject struct {
    Default interface{} `json:"default,omitempty"`
    HotPotatoRouting interface{} `json:"hot_potato_routing,omitempty"`
}