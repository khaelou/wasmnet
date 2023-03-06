// WASM Setup
if (!WebAssembly.instantiateStreaming) { // Polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await(await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject)
    }
}

const go = new Go();
let mod, inst;

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    async result => {
        mod = result.module;
        inst = result.instance;
        await go.run(inst);
    }
);
/*
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(async (result) => {
    mod = result.module
    inst = result.instance

    console.log("Instance:", inst);

    document.getElementById("runButton").disabled = false;
});

async function run() {
    await go.run(inst);
    inst = await WebAssembly.instantiate(mod, go.importObject) // reset instance
}*/