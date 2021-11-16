const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise, copyFile, interPolate, writeFile, readFile} = require('../helpers/util')

exports.command = 'resource [resourceName]'
exports.desc = 'Create new resource'
exports.builder = {
  resourceName: {
    default: '.'
  }
}

exports.handler = async function (argv) {
  
  if(argv.resourceName === '.') {
    return console.log('Resource name is required')
  }
  const resourceName = argv.resourceName
  const cwd = process.cwd()


  if(!fs.existsSync(path.join(cwd, 'main.go'))) {
    return console.log('Please execute this command from project directory')
  }

  const context = {
    resourceName,
    resourceNames:resourceName,
    resourceNameCapitalized:resourceName.charAt(0).toUpperCase() + resourceName.slice(1),
    projectName: path.basename(path.join(cwd)) 
  }
  if(!resourceName.endsWith('s')){
    context.resourceNames = resourceName + 's'
  }


  fs.ensureDirSync(path.join(cwd, resourceName))

  await copyFileWithInterpolation('controller.go', path.join(cwd,resourceName, `controller.go`), context)
  await copyFileWithInterpolation('service.go', path.join(cwd,resourceName, `service.go`), context)
  await copyFileWithInterpolation('dto.go', path.join(cwd,resourceName, `dto.go`), context)

  // update server.go file to inject newly added resource  
  let serverGoFile = await readFile(path.join(cwd,'server','server.go'))

  serverGoFile = serverGoFile.replace(/(^\s*?import\s*\([^\)]+)/gm, `$1	\"${context.projectName}/${resourceName}\"\n`)
  
  serverGoFile = serverGoFile.replace("// Routing here", `// Routing here\n	${resourceName}.Route(App)`)
  
  await writeFile(path.join(cwd,'server','server.go'), serverGoFile)


  // update db.go file to inject newly added resource  
  let dbGoFile = await readFile(path.join(cwd,'db','db.go'))

  dbGoFile = dbGoFile.replace(/(^\s*?import\s*\([^\)]+)/gm, `$1	\"${context.projectName}/${resourceName}\"\n`)
  
  dbGoFile = dbGoFile.replace("// collection set", `// collection set\n		${resourceName}.SetCollection(db)`)
  
  await writeFile(path.join(cwd,'db','db.go'), dbGoFile)


  console.log(`${resourceName} resource added`)
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