import decode from "jwt-decode";
import { navigate } from "hookrouter";

export function checkAuthorization(profile, path) {
	if (profile) {
		if (profile.exp > new Date().getTime() / 1000) {
			return true;
		}
		navigate("/login", true, { to: path });
	}
	navigate("/login", true, { to: path });
}

export function limitedAccess(roles, path) {
	const profile = GetProfile();

	if (checkAuthorization(profile, path)) {
		if (roles.length > 0) {
			for (var i in roles) {
				if (profile.role === roles[i] || profile.role === "admin")
					return;
			}
			navigate("/not-permitted");
		}
	}
}

export function GetProfile() {
	var data = null;
	try {
		data = decode(getToken());
	} catch (e) {
		return null;
	}
	return data;
}

export function getToken() {
	return localStorage.getItem("token");
}

export function isLoggedIn() {
	if (GetProfile()) {
		return true;
	} else {
		return false;
	}
}

export function logOut() {
	localStorage.removeItem("token");
}

export function GetStatus(num) {
	switch (num) {
		case 0:
			return "New";
		case 1:
			return "Accepted by Supplier";
		case 2:
			return "Accepted by Client";
		case 3:
			return "Preparation";
		case 4:
			return "Production";
		case 5:
			return "Delivered";
		case 6:
			return "Completed";
		case 7:
			return "Cancelled";
		case 8:
			return "Rejected by Supplier";
		case 9:
			return "Rejected by Client";
		default:
			break;
	}
}

export const TimestampSearch = (from, to, timestampLink) => {
	from = from.valueOf() / 1000 - 86400;
	to =
		to !== ""
			? Math.floor(to.valueOf() / 1000 + 86400)
			: new Date().valueOf() / 1000;

	timestampLink = timestampLink || "timestamp";

	return `"$and":[{"${timestampLink}":{"$gte": ${from}}},{"${timestampLink}":{"$lte": ${to}}}]`;
};

export const isEmpty = obj => {
	if (obj) {
		if (Object.entries(obj).length === 0 && obj.constructor === Object) {
			return true;
		}

		return false;
	}

	return true;
};
