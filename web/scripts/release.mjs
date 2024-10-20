import { exec } from 'child_process';
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirnameESM = path.dirname(fileURLToPath(import.meta.url));

exec('npm run build', (error, stdout, stderr) => {
  if (error) {
    console.error(`exec error: ${error}`);
    return;
  }

  if (stderr) {
    console.error(`stderr: ${stderr}`);
  }

  console.log(`build stdout: ${stdout}`);

  const curdir = __dirnameESM;
  const parentdir = path.resolve(curdir, '..');

  const distPath = path.resolve(parentdir, 'dist');
  console.info('🚀 ~ exec ~ distPath:', distPath);
  const targetPath = path.resolve(parentdir, '../server/dist');

  if (!fs.existsSync(distPath)) {
    console.error('dist not exists');
    return;
  }

  fs.rmSync(targetPath, {
    recursive: true,
    force: true,
  });

  // 将dist 文件夹复制到server 目录下

  fs.cp(distPath, targetPath, { recursive: true }, (copyError) => {
    if (copyError) {
      console.error(`copyError: ${copyError}`);
      return;
    }

    console.log('copy success');
  });
});
