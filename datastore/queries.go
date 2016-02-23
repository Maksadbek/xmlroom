package datastore

var createItemQuery string = `
	insert into items(
		id ,
		street ,
		house_number ,
		house_number_add ,
		post_code,
		city,
		sub_area ,
		estate_owner ,
		min_price ,
		project_name ,
		house_type_in_project ,
		description_nl ,
		description_fr ,
		description_en ,
		description_de ,
		description_es ,
		desctiption_it ,
		brochure_url ,
		plan_url ,
		build_year ,
		furnished ,
		hide_price ,
		update_photos ,
		parking ,
		bath ,
		separate_shower ,
		separate_toilet ,
		lift ,
		swimming_pool ,
		air_conditioning ,
		fire_place ,
		garage ,
		cellar ,
		show_house_num ,
		photos ,
		tenant ,
		area ,
		nr_of_bathrooms ,
		contract_length ,
		nr_of_rooms ,
		nr_of_living_rooms ,
		stats ,
		available ,
		insert_date ,
		type ,
		pt_quality ,
		ground_floor ,
		garden ,
		garden_loc ,
		garden_area ,
		roof_terrass ,
		roof_terrass_loc ,
		roof_terrass_area ,
		balcony ,
		balcony_loc ,
		balcony_area ,
		floor_quality ,
		rent_included 
	) 
	values (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?
	)
`
var createItemsTableQuery string = `
	create table if not exists 
	items(
		id text primary key,
		street text,
		house_number text,
		house_number_add text,
		post_code string,
		city text,
		sub_area text,
		estate_owner text,
		min_price text,
		project_name text,
		house_type_in_project text,
		description_nl text,
		description_fr text,
		description_en text,
		description_de text,
		description_es text,
		desctiption_it text,
		brochure_url text,
		plan_url text,
		build_year text,
		furnished numeric,
		hide_price numeric,
		update_photos numeric,
		parking numeric,
		bath numeric,
		separate_shower numeric,
		separate_toilet numeric,
		lift numeric,
		swimming_pool numeric,
		air_conditioning numeric,
		fire_place numeric,
		garage numeric,
		cellar numeric,
		show_house_num numeric,
		photos text,
		tenant integer,
		area integer,
		nr_of_bathrooms integer,
		contract_length integer,
		nr_of_rooms integer,
		nr_of_living_rooms integer,
		stats text,
		available numeric,
		insert_date numeric,
		type text,
		pt_quality text,
		ground_floor text,
		garden numeric,
		garden_loc text,
		garden_area integer,
		roof_terrass numeric,
		roof_terrass_loc text,
		roof_terrass_area integer,
		balcony numeric,
		balcony_loc numeric,
		balcony_area integer,
		floor_quality text,
		rent_included text	
	)
`

var createMembersTableQuery string = `
	create table if not exists members(
		id integer primary key autoincrement,
		agent_id integer
	)
`
var createMemberItemsTableQuery string = `
	create table if not exists member_items (
		id integer primary key autoincrement,
		agent_id integer,
		item_id text,
		foreign key(agent_id) references member(agent_id),
		foreign key(item_id) references items(id)
	)
`

var createMemberQuery string = `
	insert into members(
		agent_id
	) 
	values (
		?
	)
`

var createMemberItemsQuery string = `
	insert into member_items(
		agent_id,
		item_id
	) 
	values(
		?,
		?
	)
`
