/*
|--------------------------------------------------------------------------
| Routes file
|--------------------------------------------------------------------------
|
| The routes file is used for defining the HTTP routes.
|
*/

import fs from 'node:fs'

import { HttpContext } from '@adonisjs/core/http'
import Env from '#start/env'
import router from '@adonisjs/core/services/router'

const hostname = fs.readFileSync('/etc/hostname', 'utf8').trim()

router.get('/', async () => {
  return {
    result: 'Ok',
    hostname,
  }
})
router.get('/db/', async ({ response }: HttpContext) => {
  return response.notImplemented({
    result: 'Not implemented!',
    hostname,
  })
})
router.get('/chaos/', async ({ response }: HttpContext) => {
  let resp

  try {
    resp = await fetch(Env.get('CHAOS_ENDPOINT'))
  } catch (err) {
    console.error('Chaos service is unavailable.', err)
  }

  if (!resp?.ok) {
    return response.serviceUnavailable({
      result: 'Chaos service is unavailable.',
      hostname,
    })
  }

  const data = (await resp.json()) as any

  return {
    result: 'Ok',
    hostname,
    sleep_time: data.sleep_time,
  }
})
router.get('/health/', async () => {
  return {
    result: 'Ok',
    hostname,
  }
})
