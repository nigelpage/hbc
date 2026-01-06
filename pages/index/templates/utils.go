package templates

import (
	// "fmt"
	// "strings"
	// "time"
)

type TemplateIcons struct {
	LockedIcon string
	UnlockedIcon string
	HomeIcon string
	AwayIcon string
	PhoneIcon string
	AddIcon string
	HealthcheckIcon string
}

var Icons = TemplateIcons {
	LockedIcon: 	 `<svg xmlns="http://www.w3.org/2000/svg" id="lockedUnlockedIcon" class="ionicon" viewBox="0 0 512 512"><path d="M420 192h-68v-80a96 96 0 10-192 0v80H92a12 12 0 00-12 12v280a12 12 0 0012 12h328a12 12 0 0012-12V204a12 12 0 00-12-12zm-106 0H198v-80.75a58 58 0 11116 0z"/></svg>`,
	UnlockedIcon: 	 `<svg xmlns="http://www.w3.org/2000/svg" id="lockedUnlockedIcon" class="ionicon" viewBox="0 0 512 512"><path d="M420 192H198v-80.75a58.08 58.08 0 0199.07-41.07A59.4 59.4 0 01314 112h38a96 96 0 10-192 0v80H92a12 12 0 00-12 12v280a12 12 0 0012 12h328a12 12 0 0012-12V204a12 12 0 00-12-12z"/></svg>`,
	HomeIcon: 		 `<svg xmlns="http://www.w3.org/2000/svg" class="ionicon-home" viewBox="0 0 512 512"><path d="M416 174.74V48h-80v58.45L256 32 0 272h64v208h144V320h96v160h144V272h64l-96-97.26z"/></svg>`,
	AwayIcon:		 `<svg xmlns="http://www.w3.org/2000/svg" class="ionicon-away" viewBox="0 0 512 512"><path d="M488 224c-3-5-32.61-17.79-32.61-17.79 5.15-2.66 8.67-3.21 8.67-14.21 0-12-.06-16-8.06-16h-27.14c-.11-.24-.23-.49-.34-.74-17.52-38.26-19.87-47.93-46-60.95C347.47 96.88 281.76 96 256 96s-91.47.88-126.49 18.31c-26.16 13-25.51 19.69-46 60.95 0 .11-.21.4-.4.74H55.94c-7.94 0-8 4-8 16 0 11 3.52 11.55 8.67 14.21C56.61 206.21 28 220 24 224s-8 32-8 80 4 96 4 96h11.94c0 14 2.06 16 8.06 16h80c6 0 8-2 8-16h256c0 14 2 16 8 16h82c4 0 6-3 6-16h12s4-49 4-96-5-75-8-80zm-362.74 44.94A516.94 516.94 0 0170.42 272c-20.42 0-21.12 1.31-22.56-11.44a72.16 72.16 0 01.51-17.51L49 240h3c12 0 23.27.51 44.55 6.78a98 98 0 0130.09 15.06C131 265 132 268 132 268zm247.16 72L368 352H144s.39-.61-5-11.18c-4-7.82 1-12.82 8.91-15.66C163.23 319.64 208 304 256 304s93.66 13.48 108.5 21.16C370 328 376.83 330 372.42 341zm-257-136.53a96.23 96.23 0 01-9.7.07c2.61-4.64 4.06-9.81 6.61-15.21 8-17 17.15-36.24 33.44-44.35 23.54-11.72 72.33-17 110.23-17s86.69 5.24 110.23 17c16.29 8.11 25.4 27.36 33.44 44.35 2.57 5.45 4 10.66 6.68 15.33-2 .11-4.3 0-9.79-.19zm347.72 56.11C461 273 463 272 441.58 272a516.94 516.94 0 01-54.84-3.06c-2.85-.51-3.66-5.32-1.38-7.1a93.84 93.84 0 0130.09-15.06c21.28-6.27 33.26-7.11 45.09-6.69a3.22 3.22 0 013.09 3 70.18 70.18 0 01-.49 17.47z"/></svg>`,
	PhoneIcon:	 	 `<svg xmlns="http://www.w3.org/2000/svg" class="ionicon-phone" viewBox="0 0 512 512"><path d="M478.94 370.14c-5.22-5.56-23.65-22-57.53-43.75-34.13-21.94-59.3-35.62-66.52-38.81a3.83 3.83 0 00-3.92.49c-11.63 9.07-31.21 25.73-32.26 26.63-6.78 5.81-6.78 5.81-12.33 4-9.76-3.2-40.08-19.3-66.5-45.78s-43.35-57.55-46.55-67.3c-1.83-5.56-1.83-5.56 4-12.34.9-1.05 17.57-20.63 26.64-32.25a3.83 3.83 0 00.49-3.92c-3.19-7.23-16.87-32.39-38.81-66.52-21.78-33.87-38.2-52.3-43.76-57.52a3.9 3.9 0 00-3.89-.87 322.35 322.35 0 00-56 25.45A338 338 0 0033.35 92a3.83 3.83 0 00-1.26 3.74c2.09 9.74 12.08 50.4 43.08 106.72 31.63 57.48 53.55 86.93 100 133.22S252 405.21 309.54 436.84c56.32 31 97 41 106.72 43.07a3.86 3.86 0 003.75-1.26A337.73 337.73 0 00454.35 430a322.7 322.7 0 0025.45-56 3.9 3.9 0 00-.86-3.86z"/></svg>`,
	AddIcon:		 `<svg xmlns="http://www.w3.org/2000/svg" class="ionicon-add" viewBox="0 0 512 512"><path d="M256 48C141.31 48 48 141.31 48 256s93.31 208 208 208 208-93.31 208-208S370.69 48 256 48zm96 224h-80v80h-32v-80h-80v-32h80v-80h32v80h80z"/></svg>`,
	HealthcheckIcon: `<svg xmlns="http://www.w3.org/2000/svg" class="ionicon-healthcheck" viewBox="0 0 512 512"><path d="M480 272h-37.34a261.41 261.41 0 01-18.25 32H480zM32 240v32h37.34a225.1 225.1 0 01-12.4-32z"/><path d="M304 259.78l-51.73 103.46-48-160L169.89 272H69.34c10 20.92 23.5 41.41 40.63 61.68 40.12 47.46 94.25 79.75 137 108.32l9 6 9-6c42.78-28.57 96.91-60.86 137-108.32A322.78 322.78 0 00424.41 304h-98.3z"/><path d="M211.73 116.76l48 160L304 188.22 345.89 272h96.77A213.13 213.13 0 00464 176.65C463.37 114.54 413.54 64 352.92 64c-48.11 0-80.1 28-96.92 48.21C239.18 92 207.19 64 159.08 64 98.46 64 48.63 114.54 48 176.65A211.23 211.23 0 0056.94 240h93.17z"/></svg>`,
}

// const (
//   timeStoreFormat = "2006-01-02T15:04"
//   timeDisplayFormat = "Monday 02 January 2006 03:04PM"
//   timeDisplayFormatShort = "02 Jan 03:04PM"
// )

// func formatPhoneNumberURI(phoneNumber string) string {
// 	trimmed := strings.Join(strings.Fields(strings.TrimSpace(phoneNumber)), "")
// 	tel := fmt.Sprintf("href=tel:+61%s", strings.Replace(trimmed, "0", "", 1))
// 	return tel
// }

// func formatTime(timeStr string, shortTimeDisplay bool) string {
// 	t, err := time.Parse(timeStoreFormat, timeStr)
// 	if err != nil {
// 		return "Invalid time format!"
// 	}
	
// 	if shortTimeDisplay {
// 		return t.Format(timeDisplayFormatShort)
// 	}

// 	return t.Format(timeDisplayFormat)
// }
