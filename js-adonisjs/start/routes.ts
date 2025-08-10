/*
|--------------------------------------------------------------------------
| Routes file
|--------------------------------------------------------------------------
|
| The routes file is used for defining the HTTP routes.
|
*/

import { HttpContext } from '@adonisjs/core/http'
import Env from '#start/env'
import router from '@adonisjs/core/services/router'

router.get('/', async () => {
  return {
    result: 'Ok',
  }
})
router.get('/db/', async ({ response }: HttpContext) => {
  return response.notImplemented({
    result: 'Not implemented!',
  })
})
router.get('/chaos/', async ({ response }: HttpContext) => {
  const resp = await fetch(Env.get('CHAOS_ENDPOINT'))

  if (!resp.ok) {
    return response.serviceUnavailable({
      result: 'Chaos service is unavailable.',
    })
  }

  const data = (await resp.json()) as any

  return {
    result: 'Ok',
    sleep_time: data.sleep_time,
  }
})
router.get('/health/', async () => {
  return {
    result: 'Ok',
  }
})
