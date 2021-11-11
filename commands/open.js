const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise} = require('../helpers/util')

exports.command = 'open [editor]'
exports.desc = 'Open project in editor'
exports.builder = {
  editor: {
    default: '.'
  }
}

exports.handler = async function (argv) {
  if(argv.editor === 'vscode') {
    console.log(`opening ${process.cwd()} path in vscode`);
    exec('code', ['.'], {cwd: process.cwd()}, (data) => {
        console.log(data)
      }
    )
  }else{
    return console.log(`editor ${argv.editor} is not supported, only vscode is supported for now`)
  }
}
