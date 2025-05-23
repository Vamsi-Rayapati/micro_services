import { S3Client } from '@aws-sdk/client-s3';
import config from '../../../config';

const s3Client = new S3Client({
  region: 'us-east-2',
  credentials: {
    accessKeyId: config.AWS_ACCESS_KEY_ID,
    secretAccessKey: config.AWS_SECRET_ACCESS_KEY,
  },

});

export default s3Client;
