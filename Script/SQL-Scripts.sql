-- Table: public.Campaigns

-- DROP TABLE IF EXISTS public."Campaigns";

CREATE TABLE IF NOT EXISTS public."Campaigns"
(
    "ID" bigint NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "StartDate" date NOT NULL,
    "EndDate" date NOT NULL,
    "Active" boolean NOT NULL,
    CONSTRAINT "Campaigns_pkey" PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Campaigns"
    OWNER to postgres;
-----------------------------------------------------------------

-- Table: public.Events

-- DROP TABLE IF EXISTS public."Events";

CREATE TABLE IF NOT EXISTS public."Events"
(
    "ID" bigint NOT NULL,
    "CampaignId" bigint NOT NULL,
    "Event" text COLLATE pg_catalog."default" NOT NULL,
    "Point" bigint NOT NULL,
    "Description" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Events_pkey" PRIMARY KEY ("ID"),
    CONSTRAINT "Event-Frnkey" FOREIGN KEY ("CampaignId")
        REFERENCES public."Campaigns" ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Events"
    OWNER to postgres;
---------------------------------------------
-- Table: public.Scores

-- DROP TABLE IF EXISTS public."Scores";

CREATE TABLE IF NOT EXISTS public."Scores"
(
    "ID" bigint NOT NULL,
    "CampaignId" bigint NOT NULL,
    "EventId" bigint NOT NULL,
    "DateTime" date NOT NULL,
    "User" text COLLATE pg_catalog."default" NOT NULL,
    "UserEmail" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Scores_pkey" PRIMARY KEY ("ID"),
    CONSTRAINT "Score-Event-FrnKey" FOREIGN KEY ("EventId")
        REFERENCES public."Events" ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT "Score-camp-FrnKey" FOREIGN KEY ("CampaignId")
        REFERENCES public."Campaigns" ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Scores"
    OWNER to postgres;
----------------------------------------------------------

-- View: public.Campaign_View

-- DROP VIEW public."Campaign_View";

CREATE OR REPLACE VIEW public."Campaign_View"
 AS
 SELECT "Campaigns"."ID",
    "Campaigns"."Name",
    "Campaigns"."StartDate",
    "Campaigns"."EndDate",
    "Campaigns"."Active"
   FROM "Campaigns"
  WHERE "Campaigns"."StartDate" <= CURRENT_DATE AND "Campaigns"."EndDate" >= CURRENT_DATE AND "Campaigns"."Active" = true;

ALTER TABLE public."Campaign_View"
    OWNER TO postgres;

-----------------------------------------------------------------
-- View: public.Events_View

-- DROP VIEW public."Events_View";

CREATE OR REPLACE VIEW public."Events_View"
 AS
 SELECT "Events"."ID",
    "Events"."CampaignId",
    "Events"."Event",
    "Events"."Point",
    "Events"."Description"
   FROM "Events";

ALTER TABLE public."Events_View"
    OWNER TO postgres;



