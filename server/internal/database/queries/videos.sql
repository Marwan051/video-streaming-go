-- name: GetRenderedVideos :many
SELECT r.label, r.bitrate, r.path
  FROM videos v
  JOIN renditions r ON r.video_id = v.id
 WHERE v.id = :videoID
   AND v.status = 'ready'
 ORDER BY r.bitrate;
