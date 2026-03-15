import fastify from 'fastify'
import { ZodError } from 'zod'
import { balancesRoutes } from './controllers/balance/routes'
import { env } from '../../../env'

export const app = fastify()

app.register(balancesRoutes)

app.setErrorHandler((error, _, reply) => {
  if (error instanceof ZodError) {
    console.log(error.format())
    return reply
      .status(400)
      .send({ message: 'Validation error.', issues: error.format() })
  }

  if (env.NODE_ENV !== 'production') {
    console.error(error)
  } else {
    // TODO: Here we should log to an external tool like DataDog/NewRelic/Sentry
  }

  return reply.status(500).send({ message: 'Internal server error.' })
})
