#! /usr/bin/env node

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

"use strict";

const path = require("path");

globalThis.require = require;
globalThis.fs = require("fs");
globalThis.TextEncoder = require("util").TextEncoder;
globalThis.TextDecoder = require("util").TextDecoder;

globalThis.performance ??= require("performance");

globalThis.crypto ??= require("crypto");

require("../lib/wasm_exec");

const wasmPath = path.resolve(__dirname, "../wasm/main.wasm");

const go = new Go();
go.exit = process.exit;
WebAssembly.instantiate(fs.readFileSync(wasmPath), go.importObject)
  .then((result) => {
    process.on("exit", (code) => {
      // Node.js exits if no event handler is pending
      if (code === 0 && !go.exited) {
        // deadlock, make Go print error and stack traces
        go._pendingEvent = { id: 0 };
        go._resume();
      }
    });
    return go.run(result.instance);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
