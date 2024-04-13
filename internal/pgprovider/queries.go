package pgprovider

const (
	CreateUser = `INSERT INTO Users (login, password)
					VALUES (?, crypt(?, 'placeholder'))`

	CheckUser = `SELECT
					EXISTS (SELECT login FROM users WHERE login=?)`

	CheckAdmin = `SELECT
					EXISTS (SELECT login FROM admins WHERE login=?)`

	AuthUser = `SELECT (password = crypt(?, password)) 
				AS password_match
				FROM users
				WHERE login= ?;`

	AuthAdmin = `SELECT (password = crypt(?, password)) 
				AS password_match
				FROM admins
				WHERE login= ?;`

	GetBannersForTagAndFeature = `select * from banner
								where feature = ? 
								and id in 
								(select banner_id from banner_tag where tag_id = ?)`

	GetByFeature = `select * from banner where feature=?`

	GetByTag = `select * from banner 
					where id in (select banner_id from banner_tag where tag_id = ?)`

	GetAll = `select * from banner`

	UpdateBanner = `UPDATE banner
      				SET
          			feature = CASE WHEN ? = true THEN ?,
          			contents = CASE WHEN ? = true THEN ?,
					is_active = CASE WHEN ? = true THEN ?
					where id=?`

	InsertTag = `insert into banner_tag (banner_id, tag_id)
					values (?,?)
					on conflict do nothing `

	DeleteBanner = "delete from banner where id=?"
)
