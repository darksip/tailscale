// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package tsconst exports some constants used elsewhere in the
// codebase.
package tsconst

// WintunInterfaceDesc is the description attached to Tailscale
// interfaces on Windows. This is set by the WinTun driver.
const WintunInterfaceDesc = "CyberVpn Tunnel"
const WintunInterfaceDesc0_14 = "Wintun Userspace Tunnel"
