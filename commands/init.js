const fs = require('fs-extra')
const path = require('path')
const {exec, handlePromise, copyFile, interPolate, writeFile, readFile} = require('../helpers/util')

exports.command = 'init [projectName]'
exports.desc = 'Initialize new project'
exports.builder = {
  projectName: {
    default: '.'
  }
}
exports.handler = async function (argv) {
  console.log('init called for projectName', argv.projectName, process.cwd())
  const projectPath = path.join(process.cwd(),argv.projectName)
  const projectBinPath = path.join(projectPath,'bin')
  const projectGlobalPath = path.join(projectPath,'global')
  const projectServerPath = path.join(projectPath,'server')

  fs.ensureDirSync(projectBinPath)
  fs.ensureDirSync(projectGlobalPath)
  fs.ensureDirSync(projectServerPath)

  await copyFileWithInterpolation('.env',path.join(projectPath,'.env'),{projectName: argv.projectName}) 
  await copyFileWithInterpolation('main.go',path.join(projectPath,'main.go'),{projectName: argv.projectName})
  await copyFileWithInterpolation('server.go',path.join(projectServerPath,'server.go'),{projectName: argv.projectName})
  
  await copyFileTo('global.go',path.join(projectGlobalPath,'global.go'))

  await runCommand('go', ['mod', 'init', argv.projectName], projectPath)
  await runCommand('go', ['mod', 'tidy'], projectPath)

  await installAllDependencies(projectPath)
}


async function copyFileTo(srcFileName,destPath) {
  // coping db.go file
  let [fileErr, fileContent] = await handlePromise(
    copyFile(path.join(__dirname,`../templates/${srcFileName}`), destPath)
  );
  if (fileErr) {
    console.log(`Error copying ${srcFileName} template file`, fileErr)
    throw fileErr
  }
}


async function copyFileWithInterpolation(srcFileName,destPath,context){
  // coping main.go file
  let [mainFileErr, mainFileContent] = await handlePromise(
    readFile(path.join(__dirname,`../templates/${srcFileName}`))
  );
  if (mainFileErr) {
    console.log(`Error reading ${srcFileName} template file`, dbFileErr)
    throw dbFileErr
  }
  //{projectName: path.basename(projectPath)}
  //path.join(projectPath,'main.go')
  let mainFileContentInterpolated = interPolate(mainFileContent,context )
  const [mainFileErr2, mainFileContent2] = await handlePromise(writeFile(destPath, mainFileContentInterpolated))
  if (mainFileErr2) {
    console.log(`Error writing ${srcFileName} file`, mainFileErr2)
    throw mainFileErr2
  }
}


async function installAllDependencies(projectPath) {
  console.log('installing dependencies...')
  await installDependency('fiber','go',['get', 'github.com/gofiber/fiber/v2'], projectPath)
  await installDependency('godotenv', 'go',['get', 'github.com/joho/godotenv'], projectPath)
  await installDependency('mongo', 'go',['get', 'go.mongodb.org/mongo-driver/mongo'], projectPath)

}

async function installDependency(name, cmd, opts, projectPath) {
  console.log(`installing ${name}...`)
  const [error, response ] = await handlePromise(
    exec(cmd, opts, {cwd: projectPath}, (data)=>{
      console.log(data)
    })
  )
  if (error) {
    return console.log(`error while installing ${name} package: ${error.error}`)
  }else{
    console.log(`${name} installed`)
  }
}

async function runCommand( cmd, opts, projectPath) {
  const [error, response ] = await handlePromise(
    exec(cmd, opts, {cwd: projectPath}, (data)=>{
      console.log(data)
    })
  )
  if (error) {
    return console.log(`error while running command ${cmd} : ${error.error}`)
  }
}

function getDotEnvContent(projectName) {
return `
# This file is used to configure the application.
#
# The following variables are available:
#
#   - GO_ENV: app environement
GO_ENV=local
#   - PORT: the port to listen on
PORT=3005
#   - DATABASE_URL: the database url
DATABASE_URL=mongodb://localhost:27017/${projectName}
#   - DATABASE_Name: the database name
DATABASE_NAME=${projectName}

`
}

