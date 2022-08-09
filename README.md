# pokemon
consume pokemon api

filter req post format 
{
    "name": ["bulbasaur", "ivy"],
    "weight" : "<10",
    "height" : "=12",
    "ability" : [],
    "species" : []
}


request delete and add format 
{
  "ability_id" : 12,
  "pokemon_id" : 23
}


DATABASE 

table "pokemon"
CREATE TABLE public.pokemon
(
    id integer NOT NULL DEFAULT nextval('pokemon_id_seq'::regclass),
    pokemon_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    height smallint NOT NULL,
    weight smallint NOT NULL,
    "createAt" date,
    "updateAt" date,
    CONSTRAINT pokemon_pkey PRIMARY KEY (id)
)

table "species"
CREATE TABLE public.species
(
    id integer NOT NULL DEFAULT nextval('species_id_seq'::regclass),
    species_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    "createAt" date,
    "updateAt" date,
    url character varying(200) COLLATE pg_catalog."default",
    CONSTRAINT species_pkey PRIMARY KEY (id)
)

table "abilities"
CREATE TABLE public.abilities
(
    id integer NOT NULL DEFAULT nextval('abilities_id_seq'::regclass),
    ability_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    "createAt" date NOT NULL,
    "updateAt" date,
    url character varying(200) COLLATE pg_catalog."default",
    CONSTRAINT abilities_pkey PRIMARY KEY (id)
)


table "detail_abilities"
CREATE TABLE public.detail_abilities
(
    abilities_id integer NOT NULL,
    pokemon_id integer NOT NULL,
    "createAt" date,
    "updateAt" date,
    CONSTRAINT detail_abilities_abilities FOREIGN KEY (abilities_id)
        REFERENCES public.abilities (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT detail_abilities_pokemon FOREIGN KEY (pokemon_id)
        REFERENCES public.pokemon (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
        NOT VALID
)


table "detail_species"
CREATE TABLE public.detail_species
(
    pokemon_id integer NOT NULL,
    species_id integer NOT NULL,
    "createAt" date,
    "updateAt" date,
    CONSTRAINT detail_species_pokemon FOREIGN KEY (pokemon_id)
        REFERENCES public.pokemon (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT detail_species_species FOREIGN KEY (species_id)
        REFERENCES public.species (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
        NOT VALID
)
