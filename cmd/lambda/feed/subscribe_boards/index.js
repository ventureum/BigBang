exports.handler = async (event) => {
  var stream = require('getstream')
  // Get environment variables
  const { STREAM_API_KEY, STREAM_API_SECRET, APP_ID } = process.env
  const { actor, boardIds } = event
  // Instantiate a new client (server side)
  const client = await stream.connect(STREAM_API_KEY, STREAM_API_SECRET, APP_ID)
  const userFeed = await client.feed('user', actor)

  await Promise.all(boardIds.map((id) => {
    return userFeed.follow('board', id)
  }))

  const response = {
    ok: true
  }

  return response
}
