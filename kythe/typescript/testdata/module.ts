// All modules define a 'module' record representing the file as a module.
// The signature is 'module' and the path is the module path (the filename
// without an extension).
//- Mod=vname("module", _, _, "testdata/module", _).node/kind record

// The first letter in the module is tagged as defining the module.
// See discussion on emitModuleAnchor().
//- ModAnchor.node/kind anchor
//- ModAnchor./kythe/loc/start 0
//- ModAnchor./kythe/loc/end 1
//- ModAnchor defines/binding Mod

export let value = 3;

export type MyType = string;
