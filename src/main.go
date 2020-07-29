/*
   Pulse-Memberlist-Plugin - PulseHA Memberlist Plugin Event Example
   Copyright (C) 2020  Andrew Zak <andrew@pulseha.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.
   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"github.com/syleron/pulseha/src/server"
)

type PulseMemberlist bool

const PluginName = "PulseHA-Memberlist"
const PluginVersion = 1.0

/**
 * Return Plugin name
 **/
func (e PulseMemberlist) Name() string {
	return PluginName
}

/**
 * Return Plugin version
 **/
func (e PulseMemberlist) Version() float64 {
	return PluginVersion
}

/**
 * Listen for memberlist changes to perform any member logic
 */
func (e PulseMemberlist) OnMemberListStatusChange(members []server.Member) error {
	// code...
	return nil
}

var PluginGen PulseMemberlist