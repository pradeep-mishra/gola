const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise} = require('../helpers/util')

exports.command = 'test [projectName]'
exports.desc = 'Test project'
exports.builder = {
  projectName: {
    default: '.'
  }
}

exports.handler = async function (argv) {
  //console.log('run projectName', argv.projectName, process.cwd())
  if(argv.projectName === '.') {
    argv.projectName = process.cwd()
  }else{
    argv.projectName = path.join(process.cwd(), argv.projectName)
  }
  //console.log('run projectName', argv.projectName)
  const mainFileDir = argv.projectName
  const mainFilePath = path.join(mainFileDir,'main.go')
  if(fs.existsSync(mainFilePath)) {
    //console.log('executing file', mainFilePath)
    await handlePromise(exec('go', ['test', '-v', './...'], {cwd: mainFileDir}, (data) => {
        console.log(data)
      }
    ))

  }else{
    console.log('main.go not found at', mainFilePath)
  }
}
