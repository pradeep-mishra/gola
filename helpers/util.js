const { spawn } = require('child_process');
const fs = require('fs-extra');

exports.handlePromise = (promise)=>{
  if (promise && typeof promise.then === "function") {
    return promise
      .then((data) => [undefined, data])
      .catch((err) => [err, undefined]);
  }
  throw new Error("only promise is allowed");
}

exports.readFile = (filePath) => {
  return new Promise((resolve, reject) => {
    fs.readFile(filePath,{encoding:'utf8'}, (err, data) => {
      if (err) {
        reject(err);
      } else {
        resolve(data);
      }
    });
  });
}

exports.copyFile = (src, dest) => {
  return new Promise((resolve, reject) => {
    fs.copy(src, dest, (err) => {
      if (err) {
        reject(err);
      } else {
        resolve();
      }
    });
  });
}

exports.interPolate = (str, data) => {
  return str.replace(/\{\{(.*?)\}\}/g, (match, key) => { 
    return data[key] || match;
  });
}

exports.writeFile = (filePath, data) => {
  return new Promise((resolve, reject) => {
    fs.writeFile(filePath, data, (err) => {
      if (err) {
        reject(err);
      } else {
        resolve();
      }
    });
  });
}



exports.exec =  (cmd, args, options, logStream) => {
  return new Promise((resolve, reject) => {
    const opts = Object.assign({shell:true}, options);
    console.log(`executing ${cmd} ${args.join(' ')}`);
    const child = spawn(cmd, args, opts);
    let stdout = '';
    let stderr = '';

    child.stdout.on('data', data => {
      stdout += data;
      if(logStream){
        logStream(data.toString());
      }
    });

    child.stderr.on('data', data => {
      stderr += data;
      if(logStream){
        logStream(data.toString());
      }
    });

    child.on('close', code => {
      if (code !== 0) {
        reject({error : new Error(`${cmd} ${args.join(' ')} exited with code ${code}`), stderr});
      } else {
        resolve({ stdout, stderr });
      }
    });
  });
}