/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

/* eslint-disable */

var jspb = require("google-protobuf");
var goog = jspb;
var global = Function("return this")();

goog.exportSymbol("proto.Params", null, global);
goog.exportSymbol("proto.Portfolio", null, global);
goog.exportSymbol("proto.Portfolios", null, global);
goog.exportSymbol("proto.Response", null, global);
goog.exportSymbol("proto.Team", null, global);
goog.exportSymbol("proto.Teams", null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Teams = function(opt_data) {
	jspb.Message.initialize(
		this,
		opt_data,
		0,
		-1,
		proto.Teams.repeatedFields_,
		null
	);
};
goog.inherits(proto.Teams, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Teams.displayName = "proto.Teams";
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.Teams.repeatedFields_ = [1];

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Teams.prototype.toObject = function(opt_includeInstance) {
		return proto.Teams.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Teams} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Teams.toObject = function(includeInstance, msg) {
		var f,
			obj = {
				teamsList: jspb.Message.toObjectList(
					msg.getTeamsList(),
					proto.Team.toObject,
					includeInstance
				)
			};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Teams}
 */
proto.Teams.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Teams();
	return proto.Teams.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Teams} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Teams}
 */
proto.Teams.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			case 1:
				var value = new proto.Team();
				reader.readMessage(
					value,
					proto.Team.deserializeBinaryFromReader
				);
				msg.addTeams(value);
				break;
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Teams.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Teams.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Teams} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Teams.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
	f = message.getTeamsList();
	if (f.length > 0) {
		writer.writeRepeatedMessage(1, f, proto.Team.serializeBinaryToWriter);
	}
};

/**
 * repeated Team teams = 1;
 * @return {!Array<!proto.Team>}
 */
proto.Teams.prototype.getTeamsList = function() {
	return /** @type{!Array<!proto.Team>} */ (jspb.Message.getRepeatedWrapperField(
		this,
		proto.Team,
		1
	));
};

/** @param {!Array<!proto.Team>} value */
proto.Teams.prototype.setTeamsList = function(value) {
	jspb.Message.setRepeatedWrapperField(this, 1, value);
};

/**
 * @param {!proto.Team=} opt_value
 * @param {number=} opt_index
 * @return {!proto.Team}
 */
proto.Teams.prototype.addTeams = function(opt_value, opt_index) {
	return jspb.Message.addToRepeatedWrapperField(
		this,
		1,
		opt_value,
		proto.Team,
		opt_index
	);
};

proto.Teams.prototype.clearTeamsList = function() {
	this.setTeamsList([]);
};

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Team = function(opt_data) {
	jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Team, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Team.displayName = "proto.Team";
}

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Team.prototype.toObject = function(opt_includeInstance) {
		return proto.Team.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Team} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Team.toObject = function(includeInstance, msg) {
		var f,
			obj = {
				name: jspb.Message.getFieldWithDefault(msg, 1, ""),
				rating: +jspb.Message.getFieldWithDefault(msg, 2, 0.0),
				image: jspb.Message.getFieldWithDefault(msg, 3, ""),
				description: jspb.Message.getFieldWithDefault(msg, 4, ""),
				timestamp: jspb.Message.getFieldWithDefault(msg, 5, 0)
			};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Team}
 */
proto.Team.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Team();
	return proto.Team.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Team} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Team}
 */
proto.Team.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			case 1:
				var value = /** @type {string} */ (reader.readString());
				msg.setName(value);
				break;
			case 2:
				var value = /** @type {number} */ (reader.readFloat());
				msg.setRating(value);
				break;
			case 3:
				var value = /** @type {string} */ (reader.readString());
				msg.setImage(value);
				break;
			case 4:
				var value = /** @type {string} */ (reader.readString());
				msg.setDescription(value);
				break;
			case 5:
				var value = /** @type {number} */ (reader.readInt64());
				msg.setTimestamp(value);
				break;
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Team.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Team.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Team} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Team.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
	f = message.getName();
	if (f.length > 0) {
		writer.writeString(1, f);
	}
	f = message.getRating();
	if (f !== 0.0) {
		writer.writeFloat(2, f);
	}
	f = message.getImage();
	if (f.length > 0) {
		writer.writeString(3, f);
	}
	f = message.getDescription();
	if (f.length > 0) {
		writer.writeString(4, f);
	}
	f = message.getTimestamp();
	if (f !== 0) {
		writer.writeInt64(5, f);
	}
};

/**
 * optional string name = 1;
 * @return {string}
 */
proto.Team.prototype.getName = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		1,
		""
	));
};

/** @param {string} value */
proto.Team.prototype.setName = function(value) {
	jspb.Message.setProto3StringField(this, 1, value);
};

/**
 * optional float rating = 2;
 * @return {number}
 */
proto.Team.prototype.getRating = function() {
	return /** @type {number} */ (+jspb.Message.getFieldWithDefault(
		this,
		2,
		0.0
	));
};

/** @param {number} value */
proto.Team.prototype.setRating = function(value) {
	jspb.Message.setProto3FloatField(this, 2, value);
};

/**
 * optional string image = 3;
 * @return {string}
 */
proto.Team.prototype.getImage = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		3,
		""
	));
};

/** @param {string} value */
proto.Team.prototype.setImage = function(value) {
	jspb.Message.setProto3StringField(this, 3, value);
};

/**
 * optional string description = 4;
 * @return {string}
 */
proto.Team.prototype.getDescription = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		4,
		""
	));
};

/** @param {string} value */
proto.Team.prototype.setDescription = function(value) {
	jspb.Message.setProto3StringField(this, 4, value);
};

/**
 * optional int64 timestamp = 5;
 * @return {number}
 */
proto.Team.prototype.getTimestamp = function() {
	return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};

/** @param {number} value */
proto.Team.prototype.setTimestamp = function(value) {
	jspb.Message.setProto3IntField(this, 5, value);
};

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Portfolio = function(opt_data) {
	jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Portfolio, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Portfolio.displayName = "proto.Portfolio";
}

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Portfolio.prototype.toObject = function(opt_includeInstance) {
		return proto.Portfolio.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Portfolio} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Portfolio.toObject = function(includeInstance, msg) {
		var f,
			obj = {
				details: jspb.Message.getFieldWithDefault(msg, 1, ""),
				sector: jspb.Message.getFieldWithDefault(msg, 2, ""),
				image: jspb.Message.getFieldWithDefault(msg, 3, ""),
				timestamp: jspb.Message.getFieldWithDefault(msg, 4, 0),
				id: jspb.Message.getFieldWithDefault(msg, 5, "")
			};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Portfolio}
 */
proto.Portfolio.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Portfolio();
	return proto.Portfolio.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Portfolio} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Portfolio}
 */
proto.Portfolio.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			case 1:
				var value = /** @type {string} */ (reader.readString());
				msg.setDetails(value);
				break;
			case 2:
				var value = /** @type {string} */ (reader.readString());
				msg.setSector(value);
				break;
			case 3:
				var value = /** @type {string} */ (reader.readString());
				msg.setImage(value);
				break;
			case 4:
				var value = /** @type {number} */ (reader.readInt64());
				msg.setTimestamp(value);
				break;
			case 5:
				var value = /** @type {string} */ (reader.readString());
				msg.setId(value);
				break;
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Portfolio.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Portfolio.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Portfolio} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Portfolio.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
	f = message.getDetails();
	if (f.length > 0) {
		writer.writeString(1, f);
	}
	f = message.getSector();
	if (f.length > 0) {
		writer.writeString(2, f);
	}
	f = message.getImage();
	if (f.length > 0) {
		writer.writeString(3, f);
	}
	f = message.getTimestamp();
	if (f !== 0) {
		writer.writeInt64(4, f);
	}
	f = message.getId();
	if (f.length > 0) {
		writer.writeString(5, f);
	}
};

/**
 * optional string details = 1;
 * @return {string}
 */
proto.Portfolio.prototype.getDetails = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		1,
		""
	));
};

/** @param {string} value */
proto.Portfolio.prototype.setDetails = function(value) {
	jspb.Message.setProto3StringField(this, 1, value);
};

/**
 * optional string sector = 2;
 * @return {string}
 */
proto.Portfolio.prototype.getSector = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		2,
		""
	));
};

/** @param {string} value */
proto.Portfolio.prototype.setSector = function(value) {
	jspb.Message.setProto3StringField(this, 2, value);
};

/**
 * optional string image = 3;
 * @return {string}
 */
proto.Portfolio.prototype.getImage = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		3,
		""
	));
};

/** @param {string} value */
proto.Portfolio.prototype.setImage = function(value) {
	jspb.Message.setProto3StringField(this, 3, value);
};

/**
 * optional int64 timestamp = 4;
 * @return {number}
 */
proto.Portfolio.prototype.getTimestamp = function() {
	return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};

/** @param {number} value */
proto.Portfolio.prototype.setTimestamp = function(value) {
	jspb.Message.setProto3IntField(this, 4, value);
};

/**
 * optional string id = 5;
 * @return {string}
 */
proto.Portfolio.prototype.getId = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		5,
		""
	));
};

/** @param {string} value */
proto.Portfolio.prototype.setId = function(value) {
	jspb.Message.setProto3StringField(this, 5, value);
};

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Portfolios = function(opt_data) {
	jspb.Message.initialize(
		this,
		opt_data,
		0,
		-1,
		proto.Portfolios.repeatedFields_,
		null
	);
};
goog.inherits(proto.Portfolios, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Portfolios.displayName = "proto.Portfolios";
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.Portfolios.repeatedFields_ = [1];

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Portfolios.prototype.toObject = function(opt_includeInstance) {
		return proto.Portfolios.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Portfolios} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Portfolios.toObject = function(includeInstance, msg) {
		var f,
			obj = {
				portfoliosList: jspb.Message.toObjectList(
					msg.getPortfoliosList(),
					proto.Portfolio.toObject,
					includeInstance
				)
			};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Portfolios}
 */
proto.Portfolios.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Portfolios();
	return proto.Portfolios.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Portfolios} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Portfolios}
 */
proto.Portfolios.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			case 1:
				var value = new proto.Portfolio();
				reader.readMessage(
					value,
					proto.Portfolio.deserializeBinaryFromReader
				);
				msg.addPortfolios(value);
				break;
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Portfolios.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Portfolios.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Portfolios} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Portfolios.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
	f = message.getPortfoliosList();
	if (f.length > 0) {
		writer.writeRepeatedMessage(
			1,
			f,
			proto.Portfolio.serializeBinaryToWriter
		);
	}
};

/**
 * repeated Portfolio portfolios = 1;
 * @return {!Array<!proto.Portfolio>}
 */
proto.Portfolios.prototype.getPortfoliosList = function() {
	return /** @type{!Array<!proto.Portfolio>} */ (jspb.Message.getRepeatedWrapperField(
		this,
		proto.Portfolio,
		1
	));
};

/** @param {!Array<!proto.Portfolio>} value */
proto.Portfolios.prototype.setPortfoliosList = function(value) {
	jspb.Message.setRepeatedWrapperField(this, 1, value);
};

/**
 * @param {!proto.Portfolio=} opt_value
 * @param {number=} opt_index
 * @return {!proto.Portfolio}
 */
proto.Portfolios.prototype.addPortfolios = function(opt_value, opt_index) {
	return jspb.Message.addToRepeatedWrapperField(
		this,
		1,
		opt_value,
		proto.Portfolio,
		opt_index
	);
};

proto.Portfolios.prototype.clearPortfoliosList = function() {
	this.setPortfoliosList([]);
};

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Response = function(opt_data) {
	jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Response, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Response.displayName = "proto.Response";
}

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Response.prototype.toObject = function(opt_includeInstance) {
		return proto.Response.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Response} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Response.toObject = function(includeInstance, msg) {
		var f,
			obj = {};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Response}
 */
proto.Response.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Response();
	return proto.Response.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Response} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Response}
 */
proto.Response.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Response.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Response.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Response} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Response.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
};

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Params = function(opt_data) {
	jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Params, jspb.Message);
if (goog.DEBUG && !COMPILED) {
	proto.Params.displayName = "proto.Params";
}

if (jspb.Message.GENERATE_TO_OBJECT) {
	/**
	 * Creates an object representation of this proto suitable for use in Soy templates.
	 * Field names that are reserved in JavaScript and will be renamed to pb_name.
	 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
	 * For the list of reserved names please see:
	 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
	 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
	 *     for transitional soy proto support: http://goto/soy-param-migration
	 * @return {!Object}
	 */
	proto.Params.prototype.toObject = function(opt_includeInstance) {
		return proto.Params.toObject(opt_includeInstance, this);
	};

	/**
	 * Static version of the {@see toObject} method.
	 * @param {boolean|undefined} includeInstance Whether to include the JSPB
	 *     instance for transitional soy proto support:
	 *     http://goto/soy-param-migration
	 * @param {!proto.Params} msg The msg instance to transform.
	 * @return {!Object}
	 * @suppress {unusedLocalVariables} f is only used for nested messages
	 */
	proto.Params.toObject = function(includeInstance, msg) {
		var f,
			obj = {
				name: jspb.Message.getFieldWithDefault(msg, 1, ""),
				id: jspb.Message.getFieldWithDefault(msg, 2, "")
			};

		if (includeInstance) {
			obj.$jspbMessageInstance = msg;
		}
		return obj;
	};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Params}
 */
proto.Params.deserializeBinary = function(bytes) {
	var reader = new jspb.BinaryReader(bytes);
	var msg = new proto.Params();
	return proto.Params.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Params} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Params}
 */
proto.Params.deserializeBinaryFromReader = function(msg, reader) {
	while (reader.nextField()) {
		if (reader.isEndGroup()) {
			break;
		}
		var field = reader.getFieldNumber();
		switch (field) {
			case 1:
				var value = /** @type {string} */ (reader.readString());
				msg.setName(value);
				break;
			case 2:
				var value = /** @type {string} */ (reader.readString());
				msg.setId(value);
				break;
			default:
				reader.skipField();
				break;
		}
	}
	return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Params.prototype.serializeBinary = function() {
	var writer = new jspb.BinaryWriter();
	proto.Params.serializeBinaryToWriter(this, writer);
	return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Params} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Params.serializeBinaryToWriter = function(message, writer) {
	var f = undefined;
	f = message.getName();
	if (f.length > 0) {
		writer.writeString(1, f);
	}
	f = message.getId();
	if (f.length > 0) {
		writer.writeString(2, f);
	}
};

/**
 * optional string name = 1;
 * @return {string}
 */
proto.Params.prototype.getName = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		1,
		""
	));
};

/** @param {string} value */
proto.Params.prototype.setName = function(value) {
	jspb.Message.setProto3StringField(this, 1, value);
};

/**
 * optional string id = 2;
 * @return {string}
 */
proto.Params.prototype.getId = function() {
	return /** @type {string} */ (jspb.Message.getFieldWithDefault(
		this,
		2,
		""
	));
};

/** @param {string} value */
proto.Params.prototype.setId = function(value) {
	jspb.Message.setProto3StringField(this, 2, value);
};

goog.object.extend(exports, proto);