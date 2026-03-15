import { env } from '../../../env'
import { app } from './app'

app
  .listen({
    // configuração obrigatória para fazer o deploy
    host: '0.0.0.0',
    port: env.PORT,
  })
  .then(() => {
    console.log('HTTP Server Running!')
  })
