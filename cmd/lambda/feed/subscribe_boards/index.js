const Knex = require('knex')
const stream = require('getstream')

const { DB_NAME_PREFIX, STAGE, DB_USER, DB_PASSWORD, DB_HOST_POSTFIX, AUTH_LEVEL } = process.env
const DATA_BASE = DB_NAME_PREFIX + STAGE
const PORT = 5432
const SuperAuth = 'SuperAuth'
const AdminAuth = 'AdminAuth'
const UserAuth = 'UserAuth'
const NoAuth = 'NoAuth'
const Config = {
  client: 'pg',
  connection: {
    host: `${DATA_BASE}${DB_HOST_POSTFIX}`,
    port: PORT,
    user: DB_USER,
    password: DB_PASSWORD,
    database: DATA_BASE
  }
}

async function secondaryCheck (principalId, actor) {
  const knex = Knex(Config)
  const record = await knex
    .select('*')
    .from('actor_profile_records')
    .where({ 'actor': principalId, actor_profile_status: 'ACTIVATED' })
  console.log('userRecord:', record)
  knex.destroy()

  if (AUTH_LEVEL === AdminAuth && principalId !== '') {
    return record[0] !== undefined &&
      record[0].actor_type === 'ADMIN'
  } else if (AUTH_LEVEL === UserAuth && principalId !== '') {
    return (record[0] !== undefined &&
      (record[0].actor_type === 'ADMIN' || principalId === actor))
  } else {
    return (AUTH_LEVEL === NoAuth && principalId !== '')
  }
}

exports.handler = async (event) => {
  // Get environment variables
  const { STREAM_API_KEY, STREAM_API_SECRET, APP_ID } = process.env
  const { actor, boardIds } = event.body

  // perform secodary check
  if (await secondaryCheck(event.principalId, actor) === true) {
    try {
      // Instantiate a new client (server side)
      const client = await stream.connect(STREAM_API_KEY, STREAM_API_SECRET, APP_ID)
      const userFeed = await client.feed('user', actor)

      await Promise.all(boardIds.map((id) => {
        return userFeed.follow('board', id)
      }))

      return {
        ok: true
      }
    } catch (e) {
      return {
        ok: false,
        errorMessage: 'Failed to follow'
      }
    }
  }
  return {
    ok: false,
    errorMessage: 'Access denied'
  }
}
