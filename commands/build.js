const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise} = require('../helpers/util')

exports.command = 'build [projectName]'
exports.desc = 'Build project'
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
  const projectName = path.basename(mainFileDir)

  if(fs.existsSync(mainFilePath)) {
    //console.log('executing file', mainFilePath)
    await handlePromise(exec('go', ['build', '-o', `bin/.`, '.'], {cwd: mainFileDir}, (data) => {
        console.log(data)
      }
    ))

  }else{
    console.log('main.go not found at', mainFilePath)
  }
}
