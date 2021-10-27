// Kodex (Community Edition - CE) - Privacy & Security Engineering Platform
// Copyright (C) 2019-2021  KIProtect GmbH (HRB 208395B) - Germany
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package decorators

import (
	"github.com/gin-gonic/gin"
	"github.com/kiprotect/kodex"
)

func WithSettings(settings kodex.Settings) gin.HandlerFunc {

	/*
	   This decorator adds a reference to the settings object to the request context.
	*/

	decorator := func(c *gin.Context) {
		c.Set("settings", settings)
	}
	return decorator
}