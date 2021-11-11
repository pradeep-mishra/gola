const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise, copyFile, interPolate, writeFile, readFile} = require('../helpers/util')

exports.command = 'controller [controllerName]'
exports.desc = 'Create new controller'
exports.builder = {
  controllerName: {
    default: '.'
  }
}

exports.handler = async function (argv) {
  
  if(argv.controllerName === '.') {
    return console.log('Controller name is required')
  }
  const controllerName = argv.controllerName
  const cwd = process.cwd()


  if(!fs.existsSync(path.join(cwd, 'main.go'))) {
    return console.log('Please execute this command from project directory')
  }

  const context = {
    controllerName,
    controllerNames:controllerName,
    projectName: path.basename(path.join(cwd)) 
  }
  if(!controllerName.endsWith('s')){
    context.controllerNames = controllerName + 's'
  }


  fs.ensureDirSync(path.join(cwd, controllerName))

  await copyFileWithInterpolation('controller.go', path.join(cwd,controllerName, `controller.go`), context)

  let serverGoFile = await readFile(path.join(cwd,'server','server.go'))

  serverGoFile = serverGoFile.replace(/(^\s*?import\s*\([^\)]+)/gm, `$1  \"${context.projectName}/${controllerName}\"\n`)
  
  serverGoFile = serverGoFile.replace("// Routing here", `// Routing here\n	${controllerName}.Load(app)`)
  
  await writeFile(path.join(cwd,'server','server.go'), serverGoFile)

  console.log(`${controllerName} controller added`)
}


async function copyFileWithInterpolation(srcFileName,destPath,context){
  let [mainFileErr, mainFileContent] = await handlePromise(
    readFile(path.join(__dirname,`../templates/${srcFileName}`))
  );
  if (mainFileErr) {
    console.log(`Error reading ${srcFileName} template file`, dbFileErr)
    throw dbFileErr
  }

  let mainFileContentInterpolated = interPolate(mainFileContent,context )
  const [mainFileErr2, mainFileContent2] = await handlePromise(writeFile(destPath, mainFileContentInterpolated))
  if (mainFileErr2) {
    console.log(`Error writing ${srcFileName} file`, mainFileErr2)
    throw mainFileErr2
  }
}